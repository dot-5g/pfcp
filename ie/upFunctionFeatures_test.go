package ie_test

import (
	"testing"

	"github.com/dot-5g/pfcp/ie"
)

func TestGivenSerializedWhenDeserializeUPFunctionFeaturesThenDeserializedCorrectly(t *testing.T) {
	features := [](ie.UPFeature){
		ie.BUCP,
		ie.TRACE,
	}

	upFunctionFeatures, err := ie.NewUPFunctionFeatures(features)

	if err != nil {
		t.Fatalf("Error creating UPFunctionFeatures: %v", err)
	}

	serializedUPFunctionFeatures := upFunctionFeatures.Serialize()

	deserializedUPFunctionFeatures, err := ie.DeserializeUPFunctionFeatures(serializedUPFunctionFeatures)
	if err != nil {
		t.Fatalf("Error deserializing UPFunctionFeatures: %v", err)
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
