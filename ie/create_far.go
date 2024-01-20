package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type CreateFAR struct {
	Header      Header
	FARID       FARID
	ApplyAction ApplyAction
}

func NewCreateFAR(farid FARID, applyaction ApplyAction) (CreateFAR, error) {
	ieHeader := Header{
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

func (createfar CreateFAR) SetHeader(header Header) InformationElement {
	createfar.Header = header
	return createfar
}

func DeserializeCreateFAR(value []byte) (CreateFAR, error) {
	var createfar CreateFAR

	if len(value) < HeaderLength {
		return CreateFAR{}, fmt.Errorf("invalid length for CreateFAR: got %d bytes, want at least %d", len(value), HeaderLength)
	}

	buffer := bytes.NewBuffer(value)

	// Deserialize FARID
	if buffer.Len() < 2 {
		return CreateFAR{}, fmt.Errorf("not enough data for FARID type")
	}
	faridIEType := binary.BigEndian.Uint16(buffer.Next(2))

	if buffer.Len() < 2 {
		return CreateFAR{}, fmt.Errorf("not enough data for FARID length")
	}

	faridIELength := binary.BigEndian.Uint16(buffer.Next(2))
	faridIEValue := buffer.Next(int(faridIELength))

	faridIEHEader := Header{
		Type:   IEType(faridIEType),
		Length: faridIELength,
	}

	tempFarid, err := DeserializeFARID(faridIEValue)
	if err != nil {
		return CreateFAR{}, fmt.Errorf("failed to deserialize FARID: %v", err)
	}
	farid, ok := tempFarid.SetHeader(faridIEHEader).(FARID)
	if !ok {
		return CreateFAR{}, fmt.Errorf("type assertion to FarID failed")
	}

	createfar.FARID = farid

	if buffer.Len() < 2 {
		return CreateFAR{}, fmt.Errorf("not enough data for ApplyAction type")
	}
	applyactionIEType := binary.BigEndian.Uint16(buffer.Next(2))

	if buffer.Len() < 2 {
		return CreateFAR{}, fmt.Errorf("not enough data for ApplyAction length")
	}
	applyactionIELength := binary.BigEndian.Uint16(buffer.Next(2))

	if buffer.Len() < int(applyactionIELength) {
		return CreateFAR{}, fmt.Errorf("not enough data for ApplyAction value, expected %d, got %d", applyactionIELength, buffer.Len())
	}
	applyactionIEValue := buffer.Next(int(applyactionIELength))

	applyActionHeader := Header{
		Type:   IEType(applyactionIEType),
		Length: applyactionIELength,
	}

	tempApplyaction, err := DeserializeApplyAction(applyactionIEValue)
	if err != nil {
		return CreateFAR{}, fmt.Errorf("failed to deserialize ApplyAction: %v", err)
	}

	applyaction, ok := tempApplyaction.SetHeader(applyActionHeader).(ApplyAction)
	if !ok {
		return CreateFAR{}, fmt.Errorf("type assertion to ApplyAction failed")
	}

	createfar.ApplyAction = applyaction
	return createfar, nil
}
