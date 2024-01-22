package ie_test

import (
	"testing"

	"github.com/dot-5g/pfcp/ie"
)

func TestGivenCorrectParametersWhenNewCreatePDRThenFieldsSetCorrectly(t *testing.T) {
	pdrID, err := ie.NewPDRID(1)

	if err != nil {
		t.Fatalf("Error creating PDRID: %v", err)
	}

	precedence, err := ie.NewPrecedence(1)
	if err != nil {
		t.Fatalf("Error creating Precedence: %v", err)
	}

	sourceInterface, err := ie.NewSourceInterface(1)
	if err != nil {
		t.Fatalf("Error creating SourceInterface: %v", err)
	}

	sd := ie.SourceDestination{}
	ipv6DelegationBits := uint8(32)
	ueIPAddress, err := ie.NewUEIPAddress("", "", sd, ipv6DelegationBits, 0, false, true)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	pdi, err := ie.NewPDI(sourceInterface, ueIPAddress)
	if err != nil {
		t.Fatalf("Error creating PDI: %v", err)
	}

	createPDR, err := ie.NewCreatePDR(pdrID, precedence, pdi)

	if err != nil {
		t.Fatalf("Error creating CreatePDR: %v", err)
	}

	if createPDR.PDRID != pdrID {
		t.Errorf("Expected CreatePDR PDRID %v, got %v", pdrID, createPDR.PDRID)
	}

	if createPDR.Precedence != precedence {
		t.Errorf("Expected CreatePDR Precedence %v, got %v", precedence, createPDR.Precedence)
	}

	if createPDR.PDI.SourceInterface != pdi.SourceInterface {
		t.Errorf("Expected CreatePDR PDI SourceInterface %v, got %v", pdi.SourceInterface, createPDR.PDI.SourceInterface)
	}

	if createPDR.PDI.UEIPAddress.IP6PL != false {
		t.Errorf("Expected CreatePDR PDI UEIPAddress IP6PL false, got %v", createPDR.PDI.UEIPAddress.IP6PL)
	}

	if createPDR.PDI.UEIPAddress.CHV6 != true {
		t.Errorf("Expected CreatePDR PDI UEIPAddress CHV6 true, got %v", createPDR.PDI.UEIPAddress.CHV6)
	}

	if createPDR.PDI.UEIPAddress.CHV4 != false {
		t.Errorf("Expected CreatePDR PDI UEIPAddress CHV4 false, got %v", createPDR.PDI.UEIPAddress.CHV4)
	}

	if createPDR.PDI.UEIPAddress.IPv6D != true {
		t.Errorf("Expected CreatePDR PDI UEIPAddress IPv6D true, got %v", createPDR.PDI.UEIPAddress.IPv6D)
	}

	if createPDR.PDI.UEIPAddress.SD != false {
		t.Errorf("Expected CreatePDR PDI UEIPAddress SD false, got %v", createPDR.PDI.UEIPAddress.SD)
	}

	if createPDR.PDI.UEIPAddress.V4 != false {
		t.Errorf("Expected CreatePDR PDI UEIPAddress V4 false, got %v", createPDR.PDI.UEIPAddress.V4)
	}

	if createPDR.PDI.UEIPAddress.V6 != false {
		t.Errorf("Expected CreatePDR PDI UEIPAddress V6 false, got %v", createPDR.PDI.UEIPAddress.V6)
	}

	if createPDR.PDI.UEIPAddress.IPv4Address != nil {
		t.Errorf("Expected CreatePDR PDI UEIPAddress IPv4Address nil, got %v", createPDR.PDI.UEIPAddress.IPv4Address)
	}

	if createPDR.PDI.UEIPAddress.IPv6Address != nil {
		t.Errorf("Expected CreatePDR PDI UEIPAddress IPv6Address nil, got %v", createPDR.PDI.UEIPAddress.IPv6Address)
	}

	if createPDR.PDI.UEIPAddress.IPv6PrefixDelegationBits != ipv6DelegationBits {
		t.Errorf("Expected CreatePDR PDI UEIPAddress IPv6PrefixDelegationBits %d, got %d", ipv6DelegationBits, createPDR.PDI.UEIPAddress.IPv6PrefixDelegationBits)
	}

	if createPDR.PDI.UEIPAddress.IPv6PrefixLength != 0 {
		t.Errorf("Expected CreatePDR PDI UEIPAddress IPv6PrefixLength 0, got %d", createPDR.PDI.UEIPAddress.IPv6PrefixLength)
	}
}

func TestGivenSerializedWhenDeserializeCreatePDRThenFieldsSetCorrectly(t *testing.T) {
	pdrID, err := ie.NewPDRID(1)

	if err != nil {
		t.Fatalf("Error creating PDRID: %v", err)
	}

	precedence, err := ie.NewPrecedence(1)

	if err != nil {
		t.Fatalf("Error creating Precedence: %v", err)
	}

	sourceInterface, err := ie.NewSourceInterface(1)
	if err != nil {
		t.Fatalf("Error creating SourceInterface: %v", err)
	}

	sd := ie.SourceDestination{}
	ipv6DelegationBits := uint8(32)
	ueIPAddress, err := ie.NewUEIPAddress("", "", sd, ipv6DelegationBits, 0, false, true)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	pdi, err := ie.NewPDI(sourceInterface, ueIPAddress)

	if err != nil {
		t.Fatalf("Error creating PDI: %v", err)
	}

	createPDR, err := ie.NewCreatePDR(pdrID, precedence, pdi)

	if err != nil {
		t.Fatalf("Error creating CreatePDR: %v", err)
	}

	serialized := createPDR.Serialize()

	deserialized, err := ie.DeserializeCreatePDR(serialized)

	if err != nil {
		t.Fatalf("Error deserializing CreatePDR: %v", err)
	}

	if deserialized.PDRID != pdrID {
		t.Errorf("Expected CreatePDR PDRID %v, got %v", pdrID, deserialized.PDRID)
	}

	if deserialized.Precedence != precedence {
		t.Errorf("Expected CreatePDR Precedence %v, got %v", precedence, deserialized.Precedence)
	}

	if deserialized.PDI.SourceInterface != pdi.SourceInterface {
		t.Errorf("Expected CreatePDR PDI SourceInterface %v, got %v", pdi.SourceInterface, deserialized.PDI.SourceInterface)
	}

	if deserialized.PDI.UEIPAddress.IP6PL != false {
		t.Errorf("Expected CreatePDR PDI UEIPAddress IP6PL false, got %v", deserialized.PDI.UEIPAddress.IP6PL)
	}

	if deserialized.PDI.UEIPAddress.CHV6 != true {
		t.Errorf("Expected CreatePDR PDI UEIPAddress CHV6 %v, got %v", ueIPAddress.CHV6, deserialized.PDI.UEIPAddress.CHV6)
	}

	if deserialized.PDI.UEIPAddress.CHV4 != false {
		t.Errorf("Expected CreatePDR PDI UEIPAddress CHV4 false, got %v", deserialized.PDI.UEIPAddress.CHV4)
	}

	if deserialized.PDI.UEIPAddress.IPv6D != true {
		t.Errorf("Expected CreatePDR PDI UEIPAddress IPv6D true, got %v", deserialized.PDI.UEIPAddress.IPv6D)
	}

	if deserialized.PDI.UEIPAddress.SD != false {
		t.Errorf("Expected CreatePDR PDI UEIPAddress SD false, got %v", deserialized.PDI.UEIPAddress.SD)
	}

	if deserialized.PDI.UEIPAddress.V4 != false {
		t.Errorf("Expected CreatePDR PDI UEIPAddress V4 false, got %v", deserialized.PDI.UEIPAddress.V4)
	}

	if deserialized.PDI.UEIPAddress.V6 != false {
		t.Errorf("Expected CreatePDR PDI UEIPAddress V6 false, got %v", deserialized.PDI.UEIPAddress.V6)
	}

	if deserialized.PDI.UEIPAddress.IPv4Address != nil {
		t.Errorf("Expected CreatePDR PDI UEIPAddress IPv4Address nil, got %v", deserialized.PDI.UEIPAddress.IPv4Address)
	}

	if deserialized.PDI.UEIPAddress.IPv6Address != nil {
		t.Errorf("Expected CreatePDR PDI UEIPAddress IPv6Address nil, got %v", deserialized.PDI.UEIPAddress.IPv6Address)
	}

	if deserialized.PDI.UEIPAddress.IPv6PrefixDelegationBits != ipv6DelegationBits {
		t.Errorf("Expected CreatePDR PDI UEIPAddress IPv6PrefixDelegationBits %d, got %d", ipv6DelegationBits, deserialized.PDI.UEIPAddress.IPv6PrefixDelegationBits)
	}

	if deserialized.PDI.UEIPAddress.IPv6PrefixLength != 0 {
		t.Errorf("Expected CreatePDR PDI UEIPAddress IPv6PrefixLength 0, got %d", deserialized.PDI.UEIPAddress.IPv6PrefixLength)
	}

}
