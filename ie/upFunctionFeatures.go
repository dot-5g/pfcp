package ie

import (
	"bytes"
	"fmt"
)

type UPFunctionFeatures struct {
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

func NewUPFunctionFeatures(supportedFeatures []UPFeature) (UPFunctionFeatures, error) {
	featureBytes := make([]byte, 2)

	for _, feature := range supportedFeatures {
		if feature < NumberOfUPFeatures {
			byteIndex := feature / 8
			bitPosition := feature % 8
			featureBytes[byteIndex] |= 1 << bitPosition
		}
	}

	return UPFunctionFeatures{
		SupportedFeatures:            featureBytes,
		AdditionalSupportedFeatures1: nil,
		AdditionalSupportedFeatures2: nil,
	}, nil
}

func (ie UPFunctionFeatures) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Octets 5 to 6: Supported Features
	buf.Write(ie.SupportedFeatures)

	return buf.Bytes()
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

func (ie UPFunctionFeatures) GetType() IEType {
	return UPFunctionFeaturesIEType
}

func DeserializeUPFunctionFeatures(ieValue []byte) (UPFunctionFeatures, error) {
	upFuncFeatures := UPFunctionFeatures{}

	if len(ieValue) < 2 {
		return UPFunctionFeatures{}, fmt.Errorf("invalid UPFunctionFeatures length: got %d bytes, expected at least 2", len(ieValue))
	}

	upFuncFeatures.SupportedFeatures = ieValue[:2]

	if len(ieValue) > 2 {
		upFuncFeatures.AdditionalSupportedFeatures1 = ieValue[2:3]
	}
	if len(ieValue) > 3 {
		upFuncFeatures.AdditionalSupportedFeatures2 = ieValue[3:4]
	}

	return upFuncFeatures, nil
}
