package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type FAR struct {
	Header      Header
	FARID       FARID
	ApplyAction ApplyAction
}

func NewFAR(farid FARID, applyaction ApplyAction) (FAR, error) {
	ieHeader := Header{
		Type:   IEType(FARIEType),
		Length: farid.Header.Length + applyaction.Header.Length + 8,
	}

	return FAR{
		Header:      ieHeader,
		FARID:       farid,
		ApplyAction: applyaction,
	}, nil
}

func (far FAR) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 4: Header
	buf.Write(far.Header.Serialize())

	// Octets 5 to n: FAR ID
	serializedFARID := far.FARID.Serialize()
	buf.Write(serializedFARID)

	// Octets n+1 to m: Apply Action
	serializedApplyAction := far.ApplyAction.Serialize()
	buf.Write(serializedApplyAction)

	return buf.Bytes()

}

func (far FAR) IsZeroValue() bool {
	return far.Header.Length == 0
}

func DeserializeFAR(ieHeader Header, value []byte) (FAR, error) {
	var far FAR

	if len(value) < HeaderLength {
		return far, fmt.Errorf("invalid length for FAR: got %d bytes, want at least %d", len(value), HeaderLength)
	}

	far.Header = ieHeader

	buffer := bytes.NewBuffer(value)

	// Deserialize FARID
	if buffer.Len() < 2 {
		return far, fmt.Errorf("not enough data for FARID type")
	}
	faridIEType := binary.BigEndian.Uint16(buffer.Next(2))

	if buffer.Len() < 2 {
		return far, fmt.Errorf("not enough data for FARID length")
	}

	faridIELength := binary.BigEndian.Uint16(buffer.Next(2))
	faridIEValue := buffer.Next(int(faridIELength))

	faridIEHEader := Header{
		Type:   IEType(faridIEType),
		Length: faridIELength,
	}

	farid, err := DeserializeFARID(faridIEHEader, faridIEValue)
	if err != nil {
		return far, fmt.Errorf("failed to deserialize FARID: %v", err)
	}
	far.FARID = farid

	if buffer.Len() < 2 {
		return far, fmt.Errorf("not enough data for ApplyAction type")
	}
	applyactionIEType := binary.BigEndian.Uint16(buffer.Next(2))

	if buffer.Len() < 2 {
		return far, fmt.Errorf("not enough data for ApplyAction length")
	}
	applyactionIELength := binary.BigEndian.Uint16(buffer.Next(2))

	if buffer.Len() < int(applyactionIELength) {
		return far, fmt.Errorf("not enough data for ApplyAction value, expected %d, got %d", applyactionIELength, buffer.Len())
	}
	applyactionIEValue := buffer.Next(int(applyactionIELength))

	applyActionHeader := Header{
		Type:   IEType(applyactionIEType),
		Length: applyactionIELength,
	}

	applyaction, err := DeserializeApplyAction(applyActionHeader, applyactionIEValue)
	if err != nil {
		return far, fmt.Errorf("failed to deserialize ApplyAction: %v", err)
	}
	far.ApplyAction = applyaction
	return far, nil
}
