package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type CreateFAR struct {
	IEType      uint16
	Length      uint16
	FARID       FARID
	ApplyAction ApplyAction
}

func NewCreateFAR(farid FARID, applyaction ApplyAction) CreateFAR {
	return CreateFAR{
		IEType:      uint16(CreateFARIEType),
		Length:      farid.Length + applyaction.Length + 8,
		FARID:       farid,
		ApplyAction: applyaction,
	}
}

func (createfar CreateFAR) Serialize() []byte {

	buf := new(bytes.Buffer)

	// Octets 1 to 2: Type
	binary.Write(buf, binary.BigEndian, uint16(createfar.IEType))

	// Octets 3 to 4: Length
	binary.Write(buf, binary.BigEndian, uint16(createfar.Length))

	// Octets 5 to n: FAR ID
	serializedFARID := createfar.FARID.Serialize()
	buf.Write(serializedFARID)

	// Octets n+1 to m: Apply Action
	serializedApplyAction := createfar.ApplyAction.Serialize()
	buf.Write(serializedApplyAction)

	return buf.Bytes()

}

func (createfar CreateFAR) IsZeroValue() bool {
	return createfar.Length == 0
}

func DeserializeCreateFAR(ieType uint16, length uint16, value []byte) (CreateFAR, error) {
	var createfar CreateFAR

	if len(value) < IEHeaderLength {
		return createfar, fmt.Errorf("invalid length for CreateFAR: got %d bytes, want at least %d", len(value), IEHeaderLength)
	}

	createfar.IEType = ieType
	createfar.Length = length

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

	farid, err := DeserializeFARID(faridIEType, faridIELength, faridIEValue)
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

	applyaction, err := DeserializeApplyAction(applyactionIEType, applyactionIELength, applyactionIEValue)
	if err != nil {
		return createfar, fmt.Errorf("failed to deserialize ApplyAction: %v", err)
	}
	createfar.ApplyAction = applyaction
	return createfar, nil
}
