package ie

import (
	"encoding/binary"
	"fmt"
)

type UPFunctionFeatures struct {
	IEType                       uint16
	Length                       uint16
	SupportedFeatures            []byte
	AdditionalSupportedFeatures1 []byte
	AdditionalSupportedFeatures2 []byte
}

type UPFeature int

const (
	BUCP UPFeature = iota
	DDND
	DLBD
	TRST
	FTUP
	PFDM
	HEEU
	TREU
	EMPU
	PDIU
	UDBC
	QUOAC
	TRACE
	FRRT
	PFDE
	EPFAR
	DPDRA
	ADPDP
	UEIP
	SSET
	MNOP
	MTE
	BUNDL
	GCOM
	MPAS
	RTTL
	VTIME
	NORP
	IPTV
	IP6PL
	TSCU
	MPTCP
	ATSSSLL
	QFQM
	GPQM
	MTEDT
	CIOT
	ETHAR
	DDDS
	RDS
	RTTWP
	NumberOfUPFeatures
)

func NewUPFunctionFeatures(supportedFeatures []UPFeature) UPFunctionFeatures {
	featureBytes := make([]byte, 2)

	for _, feature := range supportedFeatures {
		if feature < NumberOfUPFeatures {
			byteIndex := feature / 8
			bitPosition := feature % 8
			featureBytes[byteIndex] |= 1 << bitPosition
		}
	}

	return UPFunctionFeatures{
		IEType:                       43,
		Length:                       uint16(len(featureBytes)),
		SupportedFeatures:            featureBytes,
		AdditionalSupportedFeatures1: nil,
		AdditionalSupportedFeatures2: nil,
	}
}

func (ie UPFunctionFeatures) Serialize() []byte {
	totalLength := 4 + ie.Length
	serialized := make([]byte, totalLength)

	binary.BigEndian.PutUint16(serialized[0:2], ie.IEType)
	binary.BigEndian.PutUint16(serialized[2:4], ie.Length)

	copy(serialized[4:], ie.SupportedFeatures)

	return serialized
}

func (ie UPFunctionFeatures) GetFeatures() []UPFeature {
	features := make([]UPFeature, 0)

	for i, byteValue := range ie.SupportedFeatures {
		for j := 0; j < 8; j++ {
			if byteValue&(1<<j) > 0 {
				features = append(features, UPFeature(i*8+j))
			}
		}
	}

	return features
}

func (ie UPFunctionFeatures) IsZeroValue() bool {
	return ie.Length == 0
}

func DeserializeUPFunctionFeatures(ieType uint16, ieLength uint16, ieValue []byte) (UPFunctionFeatures, error) {
	if ieType != 43 {
		return UPFunctionFeatures{}, fmt.Errorf("incorrect IE type")
	}
	if len(ieValue) != int(ieLength) {
		return UPFunctionFeatures{}, fmt.Errorf("incorrect length: expected %d, got %d", ieLength, len(ieValue))
	}

	upFuncFeatures := UPFunctionFeatures{
		IEType:                       ieType,
		Length:                       ieLength,
		SupportedFeatures:            make([]byte, 0),
		AdditionalSupportedFeatures1: make([]byte, 0),
		AdditionalSupportedFeatures2: make([]byte, 0),
	}

	if ieLength >= 1 {
		upFuncFeatures.SupportedFeatures = append(upFuncFeatures.SupportedFeatures, ieValue[0])
	}
	if ieLength >= 2 {
		upFuncFeatures.SupportedFeatures = append(upFuncFeatures.SupportedFeatures, ieValue[1])
	}

	return upFuncFeatures, nil
}
