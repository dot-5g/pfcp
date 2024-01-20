package ie_test

import (
	"net"
	"testing"

	"github.com/dot-5g/pfcp/ie"
)

func TestGivenIPv4AddressWhenNewUEIPAddressThenFieldsSetCorrectly(t *testing.T) {
	sd := ie.SourceDestination{}
	ipv4Address := "1.2.3.4"
	ueIPAddress, err := ie.NewUEIPAddress(ipv4Address, "", sd, 0, 0, false, false)

	if err != nil {
		t.Fatalf("Error creating UEIPAddress: %v", err)
	}

	if ueIPAddress.Header.Type != 93 {
		t.Errorf("Expected UEIPAddress, got %d", ueIPAddress.Header.Type)
	}

	if ueIPAddress.Header.Length != 5 {
		t.Errorf("Expected UEIPAddress length 5, got %d", ueIPAddress.Header.Length)
	}

	if ueIPAddress.IP6PL != false {
		t.Errorf("Expected UEIPAddress IP6PL false, got %v", ueIPAddress.IP6PL)
	}

	if ueIPAddress.CHV6 != false {
		t.Errorf("Expected UEIPAddress CHV6 false, got %v", ueIPAddress.CHV6)
	}

	if ueIPAddress.CHV4 != false {
		t.Errorf("Expected UEIPAddress CHV4 false, got %v", ueIPAddress.CHV4)
	}

	if ueIPAddress.IPv6D != false {
		t.Errorf("Expected UEIPAddress IPv6D false, got %v", ueIPAddress.IPv6D)
	}

	if ueIPAddress.SD != false {
		t.Errorf("Expected UEIPAddress SD false, got %v", ueIPAddress.SD)
	}

	if ueIPAddress.V4 != true {
		t.Errorf("Expected UEIPAddress IPv4 true, got %v", ueIPAddress.V4)
	}

	if ueIPAddress.V6 != false {
		t.Errorf("Expected UEIPAddress IPv6 false, got %v", ueIPAddress.V6)
	}

	if net.IP(ueIPAddress.IPv4Address).String() != ipv4Address {
		t.Errorf("Expected UEIPAddress IPv4 address %s, got %s", ipv4Address, net.IP(ueIPAddress.IPv4Address).String())
	}

	if ueIPAddress.IPv6Address != nil {
		t.Errorf("Expected UEIPAddress IPv6 address nil, got %s", net.IP(ueIPAddress.IPv6Address).String())
	}

	if ueIPAddress.IPv6PrefixDelegationBits != 0 {
		t.Errorf("Expected UEIPAddress IPv6 prefix delegation bits 0, got %d", ueIPAddress.IPv6PrefixDelegationBits)
	}

	if ueIPAddress.IPv6PrefixLength != 0 {
		t.Errorf("Expected UEIPAddress IPv6 prefix length 0, got %d", ueIPAddress.IPv6PrefixLength)
	}

}

func TestGivenIPv6AddresWhenNewUEIPAddressThenFieldsSetCorrectly(t *testing.T) {
	sd := ie.SourceDestination{}
	ipv6Address := "2001:db8::1"
	ueIPAddress, err := ie.NewUEIPAddress("", ipv6Address, sd, 0, 0, false, false)
	if err != nil {
		t.Fatalf("Error creating UEIPAddress: %v", err)
	}

	if ueIPAddress.Header.Type != 93 {
		t.Errorf("Expected UEIPAddress, got %d", ueIPAddress.Header.Type)
	}

	if ueIPAddress.Header.Length != 17 {
		t.Errorf("Expected UEIPAddress length 17, got %d", ueIPAddress.Header.Length)
	}

	if ueIPAddress.IP6PL != false {
		t.Errorf("Expected UEIPAddress IP6PL false, got %v", ueIPAddress.IP6PL)
	}

	if ueIPAddress.CHV6 != false {
		t.Errorf("Expected UEIPAddress CHV6 false, got %v", ueIPAddress.CHV6)
	}

	if ueIPAddress.CHV4 != false {
		t.Errorf("Expected UEIPAddress CHV4 false, got %v", ueIPAddress.CHV4)
	}

	if ueIPAddress.IPv6D != false {
		t.Errorf("Expected UEIPAddress IPv6D false, got %v", ueIPAddress.IPv6D)
	}

	if ueIPAddress.SD != false {
		t.Errorf("Expected UEIPAddress SD false, got %v", ueIPAddress.SD)
	}

	if ueIPAddress.V4 != false {
		t.Errorf("Expected UEIPAddress IPv4 false, got %v", ueIPAddress.V4)
	}

	if ueIPAddress.V6 != true {
		t.Errorf("Expected UEIPAddress IPv6 true, got %v", ueIPAddress.V6)
	}

	if ueIPAddress.IPv4Address != nil {
		t.Errorf("Expected UEIPAddress IPv4 address nil, got %s", net.IP(ueIPAddress.IPv4Address).String())
	}

	if net.IP(ueIPAddress.IPv6Address).String() != ipv6Address {
		t.Errorf("Expected UEIPAddress IPv6 address %s, got %s", ipv6Address, net.IP(ueIPAddress.IPv6Address).String())
	}

	if ueIPAddress.IPv6PrefixDelegationBits != 0 {
		t.Errorf("Expected UEIPAddress IPv6 prefix delegation bits 0, got %d", ueIPAddress.IPv6PrefixDelegationBits)
	}

	if ueIPAddress.IPv6PrefixLength != 0 {
		t.Errorf("Expected UEIPAddress IPv6 prefix length 0, got %d", ueIPAddress.IPv6PrefixLength)
	}

}

func TestGivenChooseV4WhenNewUEIPAddressThenFieldsSetCorrectly(t *testing.T) {
	sd := ie.SourceDestination{}
	ueIPAddress, err := ie.NewUEIPAddress("", "", sd, 0, 0, true, false)
	if err != nil {
		t.Fatalf("Error creating UEIPAddress: %v", err)
	}

	if ueIPAddress.Header.Type != 93 {
		t.Errorf("Expected UEIPAddress, got %d", ueIPAddress.Header.Type)
	}

	if ueIPAddress.Header.Length != 1 {
		t.Errorf("Expected UEIPAddress length 1, got %d", ueIPAddress.Header.Length)
	}

	if ueIPAddress.IP6PL != false {
		t.Errorf("Expected UEIPAddress IP6PL false, got %v", ueIPAddress.IP6PL)
	}

	if ueIPAddress.CHV6 != false {
		t.Errorf("Expected UEIPAddress CHV6 false, got %v", ueIPAddress.CHV6)
	}

	if ueIPAddress.CHV4 != true {
		t.Errorf("Expected UEIPAddress CHV4 true, got %v", ueIPAddress.CHV4)
	}

	if ueIPAddress.IPv6D != false {
		t.Errorf("Expected UEIPAddress IPv6D false, got %v", ueIPAddress.IPv6D)
	}

	if ueIPAddress.SD != false {
		t.Errorf("Expected UEIPAddress SD false, got %v", ueIPAddress.SD)
	}

	if ueIPAddress.V4 != false {
		t.Errorf("Expected UEIPAddress IPv4 false, got %v", ueIPAddress.V4)
	}

	if ueIPAddress.V6 != false {
		t.Errorf("Expected UEIPAddress IPv6 false, got %v", ueIPAddress.V6)
	}

	if ueIPAddress.IPv4Address != nil {
		t.Errorf("Expected UEIPAddress IPv4 address nil, got %s", net.IP(ueIPAddress.IPv4Address).String())
	}

	if ueIPAddress.IPv6Address != nil {
		t.Errorf("Expected UEIPAddress IPv6 address nil, got %s", net.IP(ueIPAddress.IPv6Address).String())
	}

	if ueIPAddress.IPv6PrefixDelegationBits != 0 {
		t.Errorf("Expected UEIPAddress IPv6 prefix delegation bits 0, got %d", ueIPAddress.IPv6PrefixDelegationBits)
	}

	if ueIPAddress.IPv6PrefixLength != 0 {
		t.Errorf("Expected UEIPAddress IPv6 prefix length 0, got %d", ueIPAddress.IPv6PrefixLength)
	}

}

func TestGivenChooseV6WithDelegationWhenNewUEIPAddressThenFieldsSetCorrectly(t *testing.T) {
	sd := ie.SourceDestination{}
	ipv6PrefixDelegationBits := uint8(32)
	ueIPAddress, err := ie.NewUEIPAddress("", "", sd, ipv6PrefixDelegationBits, 0, false, true)
	if err != nil {
		t.Fatalf("Error creating UEIPAddress: %v", err)
	}

	if ueIPAddress.Header.Type != 93 {
		t.Errorf("Expected UEIPAddress, got %d", ueIPAddress.Header.Type)
	}

	if ueIPAddress.Header.Length != 2 {
		t.Errorf("Expected UEIPAddress length 2, got %d", ueIPAddress.Header.Length)
	}

	if ueIPAddress.IP6PL != false {
		t.Errorf("Expected UEIPAddress IP6PL false, got %v", ueIPAddress.IP6PL)
	}

	if ueIPAddress.CHV6 != true {
		t.Errorf("Expected UEIPAddress CHV6 true, got %v", ueIPAddress.CHV6)
	}

	if ueIPAddress.CHV4 != false {
		t.Errorf("Expected UEIPAddress CHV4 false, got %v", ueIPAddress.CHV4)
	}

	if ueIPAddress.IPv6D != true {
		t.Errorf("Expected UEIPAddress IPv6D true, got %v", ueIPAddress.IPv6D)
	}

	if ueIPAddress.SD != false {
		t.Errorf("Expected UEIPAddress SD false, got %v", ueIPAddress.SD)
	}

	if ueIPAddress.V4 != false {
		t.Errorf("Expected UEIPAddress IPv4 false, got %v", ueIPAddress.V4)
	}

	if ueIPAddress.V6 != false {
		t.Errorf("Expected UEIPAddress IPv6 false, got %v", ueIPAddress.V6)
	}

	if ueIPAddress.IPv4Address != nil {
		t.Errorf("Expected UEIPAddress IPv4 address nil, got %s", net.IP(ueIPAddress.IPv4Address).String())
	}

	if ueIPAddress.IPv6Address != nil {
		t.Errorf("Expected UEIPAddress IPv6 address nil, got %s", net.IP(ueIPAddress.IPv6Address).String())
	}

	if ueIPAddress.IPv6PrefixDelegationBits != ipv6PrefixDelegationBits {
		t.Errorf("Expected UEIPAddress IPv6 prefix delegation bits %d, got %d", ipv6PrefixDelegationBits, ueIPAddress.IPv6PrefixDelegationBits)
	}

	if ueIPAddress.IPv6PrefixLength != 0 {
		t.Errorf("Expected UEIPAddress IPv6 prefix length 0, got %d", ueIPAddress.IPv6PrefixLength)
	}

}

func TestGivenChooseV6WithPrefixLengthWhenNewUEIPAddressThenFieldsSetCorrectly(t *testing.T) {
	sd := ie.SourceDestination{}
	prefixLength := uint8(32)
	ueIPAddress, err := ie.NewUEIPAddress("", "", sd, 0, prefixLength, false, true)
	if err != nil {
		t.Fatalf("Error creating UEIPAddress: %v", err)
	}

	if ueIPAddress.Header.Type != 93 {
		t.Errorf("Expected UEIPAddress, got %d", ueIPAddress.Header.Type)
	}

	if ueIPAddress.Header.Length != 2 {
		t.Errorf("Expected UEIPAddress length 2, got %d", ueIPAddress.Header.Length)
	}

	if ueIPAddress.IP6PL != true {
		t.Errorf("Expected UEIPAddress IP6PL true, got %v", ueIPAddress.IP6PL)
	}

	if ueIPAddress.CHV6 != true {
		t.Errorf("Expected UEIPAddress CHV6 true, got %v", ueIPAddress.CHV6)
	}

	if ueIPAddress.CHV4 != false {
		t.Errorf("Expected UEIPAddress CHV4 false, got %v", ueIPAddress.CHV4)
	}

	if ueIPAddress.IPv6D != false {
		t.Errorf("Expected UEIPAddress IPv6D false, got %v", ueIPAddress.IPv6D)
	}

	if ueIPAddress.SD != false {
		t.Errorf("Expected UEIPAddress SD false, got %v", ueIPAddress.SD)
	}

	if ueIPAddress.V4 != false {
		t.Errorf("Expected UEIPAddress IPv4 false, got %v", ueIPAddress.V4)
	}

	if ueIPAddress.V6 != false {
		t.Errorf("Expected UEIPAddress IPv6 false, got %v", ueIPAddress.V6)
	}

	if ueIPAddress.IPv4Address != nil {
		t.Errorf("Expected UEIPAddress IPv4 address nil, got %s", net.IP(ueIPAddress.IPv4Address).String())
	}

	if ueIPAddress.IPv6Address != nil {
		t.Errorf("Expected UEIPAddress IPv6 address nil, got %s", net.IP(ueIPAddress.IPv6Address).String())
	}

	if ueIPAddress.IPv6PrefixDelegationBits != 0 {
		t.Errorf("Expected UEIPAddress IPv6 prefix delegation bits 0, got %d", ueIPAddress.IPv6PrefixDelegationBits)
	}

	if ueIPAddress.IPv6PrefixLength != prefixLength {
		t.Errorf("Expected UEIPAddress IPv6 prefix length %d, got %d", prefixLength, ueIPAddress.IPv6PrefixLength)
	}

}

func TestGivenSerializesWhenDeserializeUEIPAddressThenFieldsSetCorrectly(t *testing.T) {
	sd := ie.SourceDestination{}
	prefixLength := uint8(32)
	ueIPAddress, err := ie.NewUEIPAddress("", "", sd, 0, prefixLength, false, true)
	if err != nil {
		t.Fatalf("Error creating UEIPAddress: %v", err)
	}

	serializedUEIPAddress := ueIPAddress.Serialize()

	ieHeader := ie.Header{
		Type:   93,
		Length: 2,
	}

	deserialized, err := ie.DeserializeUEIPAddress(ieHeader, serializedUEIPAddress[4:])

	if err != nil {
		t.Fatalf("Error deserializing UEIPAddress: %v", err)
	}

	if deserialized.Header.Type != 93 {
		t.Errorf("Expected UEIPAddress, got %d", deserialized.Header.Type)
	}

	if deserialized.Header.Length != 2 {
		t.Errorf("Expected UEIPAddress length 2, got %d", deserialized.Header.Length)
	}

	if deserialized.IP6PL != true {
		t.Errorf("Expected UEIPAddress IP6PL true, got %v", deserialized.IP6PL)
	}

	if deserialized.CHV6 != true {
		t.Errorf("Expected UEIPAddress CHV6 true, got %v", deserialized.CHV6)
	}

	if deserialized.CHV4 != false {
		t.Errorf("Expected UEIPAddress CHV4 false, got %v", deserialized.CHV4)
	}

	if deserialized.IPv6D != false {
		t.Errorf("Expected UEIPAddress IPv6D false, got %v", deserialized.IPv6D)
	}

	if deserialized.SD != false {
		t.Errorf("Expected UEIPAddress SD false, got %v", deserialized.SD)
	}

	if deserialized.V4 != false {
		t.Errorf("Expected UEIPAddress IPv4 false, got %v", deserialized.V4)
	}

	if deserialized.V6 != false {
		t.Errorf("Expected UEIPAddress IPv6 false, got %v", deserialized.V6)
	}

	if deserialized.IPv4Address != nil {
		t.Errorf("Expected UEIPAddress IPv4 address nil, got %s", net.IP(deserialized.IPv4Address).String())
	}

	if deserialized.IPv6Address != nil {
		t.Errorf("Expected UEIPAddress IPv6 address nil, got %s", net.IP(deserialized.IPv6Address).String())
	}

	if deserialized.IPv6PrefixDelegationBits != 0 {
		t.Errorf("Expected UEIPAddress IPv6 prefix delegation bits 0, got %d", deserialized.IPv6PrefixDelegationBits)
	}

	if deserialized.IPv6PrefixLength != prefixLength {
		t.Errorf("Expected UEIPAddress IPv6 prefix length %d, got %d", prefixLength, deserialized.IPv6PrefixLength)
	}

}
