package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type CreateFAR struct {
	Header      IEHeader
	FARID       FARID
	ApplyAction ApplyAction
}

func NewCreateFAR(farid FARID, applyaction ApplyAction) (CreateFAR, error) {
	ieHeader := IEHeader{
		Type:   IEType(CreateFARIEType),
		Length: farid.Header.Length + applyaction.Header.Length + 8,
	}

	return CreateFAR{
		Header:      ieHeader,
		FARID:       farid,
		ApplyAction: applyaction,
	}, nil
}

func (createfar CreateFAR) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 1 to 4: Header
	buf.Write(createfar.Header.Serialize())

	// Octets 5 to n: FAR ID
	serializedFARID := createfar.FARID.Serialize()
	buf.Write(serializedFARID)

	// Octets n+1 to m: Apply Action
	serializedApplyAction := createfar.ApplyAction.Serialize()
	buf.Write(serializedApplyAction)

	return buf.Bytes()

}

func (createfar CreateFAR) IsZeroValue() bool {
	return createfar.Header.Length == 0
}

func DeserializeCreateFAR(ieHeader IEHeader, value []byte) (CreateFAR, error) {
	var createfar CreateFAR

	if len(value) < HeaderLength {
		return createfar, fmt.Errorf("invalid length for CreateFAR: got %d bytes, want at least %d", len(value), HeaderLength)
	}

	createfar.Header = ieHeader

	buffer := bytes.NewBuffer(value)

	// Deserialize FARID
	if buffer.Len() < 2 {
		return createfar, fmt.Errorf("not enough data for FARID type")
	}
	faridIEType := binary.BigEndian.Uint16(buffer.Next(2))

	if buffer.Len() < 2 {
		return createfar, fmt.Errorf("not enough data for FARID length")
	}

	faridIELength := binary.BigEndian.Uint16(buffer.Next(2))
	faridIEValue := buffer.Next(int(faridIELength))

	faridIEHEader := IEHeader{
		Type:   IEType(faridIEType),
		Length: faridIELength,
	}

	farid, err := DeserializeFARID(faridIEHEader, faridIEValue)
	if err != nil {
		return createfar, fmt.Errorf("failed to deserialize FARID: %v", err)
	}
	createfar.FARID = farid

	if buffer.Len() < 2 {
		return createfar, fmt.Errorf("not enough data for ApplyAction type")
	}
	applyactionIEType := binary.BigEndian.Uint16(buffer.Next(2))

	if buffer.Len() < 2 {
		return createfar, fmt.Errorf("not enough data for ApplyAction length")
	}
	applyactionIELength := binary.BigEndian.Uint16(buffer.Next(2))

	if buffer.Len() < int(applyactionIELength) {
		return createfar, fmt.Errorf("not enough data for ApplyAction value, expected %d, got %d", applyactionIELength, buffer.Len())
	}
	applyactionIEValue := buffer.Next(int(applyactionIELength))

	applyActionIEHeader := IEHeader{
		Type:   IEType(applyactionIEType),
		Length: applyactionIELength,
	}

	applyaction, err := DeserializeApplyAction(applyActionIEHeader, applyactionIEValue)
	if err != nil {
		return createfar, fmt.Errorf("failed to deserialize ApplyAction: %v", err)
	}
	createfar.ApplyAction = applyaction
	return createfar, nil
}
