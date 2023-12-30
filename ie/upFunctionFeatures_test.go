package ie_test

import (
	"testing"

	"github.com/dot-5g/pfcp/ie"
)

func TestGivenSerializedWhenDeserializedThenDeserializedCorrectly(t *testing.T) {
	features := [](ie.UPFeature){
		ie.BUCP,
		ie.TRACE,
	}

	upFunctionFeatures := ie.NewUPFunctionFeatures(features)
	serializedUPFunctionFeatures := upFunctionFeatures.Serialize()

	deserializedUPFunctionFeatures, err := ie.DeserializeUPFunctionFeatures(43, 2, serializedUPFunctionFeatures[4:])
	if err != nil {
		t.Fatalf("Error deserializing UPFunctionFeatures: %v", err)
	}

	if deserializedUPFunctionFeatures.IEType != 43 {
		t.Errorf("Expected IE type 43, got %d", deserializedUPFunctionFeatures.IEType)
	}

	if deserializedUPFunctionFeatures.Length != 2 {
		t.Errorf("Expected IE length 2, got %d", deserializedUPFunctionFeatures.Length)
	}

	if len(deserializedUPFunctionFeatures.SupportedFeatures) != 2 {
		t.Errorf("Expected 2 supported features, got %d", len(deserializedUPFunctionFeatures.SupportedFeatures))
	}

	deserializedFeatures := deserializedUPFunctionFeatures.GetFeatures()

	if len(deserializedFeatures) != 2 {
		t.Errorf("Expected 2 features, got %d", len(deserializedFeatures))
	}

	if deserializedFeatures[0] != ie.BUCP {
		t.Errorf("Expected BUCP feature, got %v", deserializedFeatures[0])
	}

	if deserializedFeatures[1] != ie.TRACE {
		t.Errorf("Expected TRACE feature, got %v", deserializedFeatures[1])
	}

}
