package ie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type CreateFAR struct {
	FARID       FARID
	ApplyAction ApplyAction
}

func NewCreateFAR(farid FARID, applyaction ApplyAction) (CreateFAR, error) {

	return CreateFAR{
		FARID:       farid,
		ApplyAction: applyaction,
	}, nil
}

func (createfar CreateFAR) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 5 to n: FAR ID
	serializedFARID := createfar.FARID.Serialize()
	farIDLength := uint16(len(serializedFARID))
	farIDHeader := Header{
		Type:   createfar.FARID.GetType(),
		Length: farIDLength,
	}
	buf.Write(farIDHeader.Serialize())
	buf.Write(serializedFARID)

	// Octets n+1 to m: Apply Action
	serializedApplyAction := createfar.ApplyAction.Serialize()
	applyActionLength := uint16(len(serializedApplyAction))
	applyActionHeader := Header{
		Type:   createfar.ApplyAction.GetType(),
		Length: applyActionLength,
	}
	buf.Write(applyActionHeader.Serialize())
	buf.Write(serializedApplyAction)

	return buf.Bytes()

}

func (createfar CreateFAR) GetType() IEType {
	return CreateFARIEType
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
	buffer.Next(2)

	if buffer.Len() < 2 {
		return CreateFAR{}, fmt.Errorf("not enough data for FARID length")
	}

	faridIELength := binary.BigEndian.Uint16(buffer.Next(2))
	faridIEValue := buffer.Next(int(faridIELength))

	farid, err := DeserializeFARID(faridIEValue)
	if err != nil {
		return CreateFAR{}, fmt.Errorf("failed to deserialize FARID: %v", err)
	}

	createfar.FARID = farid

	if buffer.Len() < 2 {
		return CreateFAR{}, fmt.Errorf("not enough data for ApplyAction type")
	}

	buffer.Next(2)

	if buffer.Len() < 2 {
		return CreateFAR{}, fmt.Errorf("not enough data for ApplyAction length")
	}
	applyactionIELength := binary.BigEndian.Uint16(buffer.Next(2))

	if buffer.Len() < int(applyactionIELength) {
		return CreateFAR{}, fmt.Errorf("not enough data for ApplyAction value, expected %d, got %d", applyactionIELength, buffer.Len())
	}
	applyactionIEValue := buffer.Next(int(applyactionIELength))

	applyaction, err := DeserializeApplyAction(applyactionIEValue)
	if err != nil {
		return CreateFAR{}, fmt.Errorf("failed to deserialize ApplyAction: %v", err)
	}

	createfar.ApplyAction = applyaction
	return createfar, nil
}
