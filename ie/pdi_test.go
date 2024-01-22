package ie_test

import (
	"testing"

	"github.com/dot-5g/pfcp/ie"
)

func TestGivenCorrectPDIWhenNewPDIThenFieldsSetCorrectly(t *testing.T) {
	sourceInterface, err := ie.NewSourceInterface(4)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	sd := ie.SourceDestination{}
	prefixLength := uint8(32)
	ueIPAddress, err := ie.NewUEIPAddress("", "", sd, 0, prefixLength, false, true)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	pdi, err := ie.NewPDI(sourceInterface, ueIPAddress)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if pdi.SourceInterface != sourceInterface {
		t.Errorf("Expected SourceInterface %v, got %v", sourceInterface, pdi.SourceInterface)
	}

	if pdi.UEIPAddress.IP6PL != true {
		t.Errorf("Expected UEIPAddress IP6PL true, got %v", pdi.UEIPAddress.IP6PL)
	}

	if pdi.UEIPAddress.CHV6 != true {
		t.Errorf("Expected UEIPAddress CHV6 true, got %v", pdi.UEIPAddress.CHV6)
	}

	if pdi.UEIPAddress.CHV4 != false {
		t.Errorf("Expected UEIPAddress CHV4 false, got %v", pdi.UEIPAddress.CHV4)
	}

	if pdi.UEIPAddress.IPv6D != false {
		t.Errorf("Expected UEIPAddress IPv6D false, got %v", pdi.UEIPAddress.IPv6D)
	}

	if pdi.UEIPAddress.SD != false {
		t.Errorf("Expected UEIPAddress SD false, got %v", pdi.UEIPAddress.SD)
	}

	if pdi.UEIPAddress.V4 != false {
		t.Errorf("Expected UEIPAddress V4 false, got %v", pdi.UEIPAddress.V4)
	}

	if pdi.UEIPAddress.V6 != false {
		t.Errorf("Expected UEIPAddress V6 false, got %v", pdi.UEIPAddress.V6)
	}

	if pdi.UEIPAddress.IPv4Address != nil {
		t.Errorf("Expected UEIPAddress IPv4Address nil, got %v", pdi.UEIPAddress.IPv4Address)
	}

	if pdi.UEIPAddress.IPv6Address != nil {
		t.Errorf("Expected UEIPAddress IPv6Address nil, got %v", pdi.UEIPAddress.IPv6Address)
	}

	if pdi.UEIPAddress.IPv6PrefixDelegationBits != 0 {
		t.Errorf("Expected UEIPAddress IPv6PrefixDelegationBits 0, got %d", pdi.UEIPAddress.IPv6PrefixDelegationBits)
	}

	if pdi.UEIPAddress.IPv6PrefixLength != prefixLength {
		t.Errorf("Expected UEIPAddress IPv6PrefixLength %d, got %d", prefixLength, pdi.UEIPAddress.IPv6PrefixLength)
	}
}

func TestGivenPDISerializedWhenDeserializeThenFieldsSetCorrectly(t *testing.T) {
	sourceInterface, err := ie.NewSourceInterface(4)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	sd := ie.SourceDestination{}
	prefixLength := uint8(32)
	ueIPAddress, err := ie.NewUEIPAddress("", "", sd, 0, prefixLength, false, true)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	pdi, err := ie.NewPDI(sourceInterface, ueIPAddress)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	pdiSerialized := pdi.Serialize()

	deserializedPDI, err := ie.DeserializePDI(pdiSerialized)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if deserializedPDI.SourceInterface != sourceInterface {
		t.Errorf("Expected SourceInterface %v, got %v", sourceInterface, deserializedPDI.SourceInterface)
	}

	if deserializedPDI.UEIPAddress.IP6PL != true {
		t.Errorf("Expected UEIPAddress IP6PL true, got %v", deserializedPDI.UEIPAddress.IP6PL)
	}

	if deserializedPDI.UEIPAddress.CHV6 != true {
		t.Errorf("Expected UEIPAddress CHV6 true, got %v", deserializedPDI.UEIPAddress.CHV6)
	}

	if deserializedPDI.UEIPAddress.CHV4 != false {
		t.Errorf("Expected UEIPAddress CHV4 false, got %v", deserializedPDI.UEIPAddress.CHV4)
	}

	if deserializedPDI.UEIPAddress.IPv6D != false {
		t.Errorf("Expected UEIPAddress IPv6D false, got %v", deserializedPDI.UEIPAddress.IPv6D)
	}

	if deserializedPDI.UEIPAddress.SD != false {
		t.Errorf("Expected UEIPAddress SD false, got %v", deserializedPDI.UEIPAddress.SD)
	}

	if deserializedPDI.UEIPAddress.V4 != false {
		t.Errorf("Expected UEIPAddress V4 false, got %v", deserializedPDI.UEIPAddress.V4)
	}

	if deserializedPDI.UEIPAddress.V6 != false {
		t.Errorf("Expected UEIPAddress V6 false, got %v", deserializedPDI.UEIPAddress.V6)
	}

	if deserializedPDI.UEIPAddress.IPv4Address != nil {
		t.Errorf("Expected UEIPAddress IPv4Address nil, got %v", deserializedPDI.UEIPAddress.IPv4Address)
	}

	if deserializedPDI.UEIPAddress.IPv6Address != nil {
		t.Errorf("Expected UEIPAddress IPv6Address nil, got %v", deserializedPDI.UEIPAddress.IPv6Address)
	}

	if deserializedPDI.UEIPAddress.IPv6PrefixDelegationBits != 0 {
		t.Errorf("Expected UEIPAddress IPv6PrefixDelegationBits 0, got %d", deserializedPDI.UEIPAddress.IPv6PrefixDelegationBits)
	}

	if deserializedPDI.UEIPAddress.IPv6PrefixLength != prefixLength {
		t.Errorf("Expected UEIPAddress IPv6PrefixLength %d, got %d", prefixLength, deserializedPDI.UEIPAddress.IPv6PrefixLength)
	}
}
