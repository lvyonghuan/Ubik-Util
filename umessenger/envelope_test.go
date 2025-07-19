package umessenger

import "testing"

func TestUnmarshalEnvelopeValidData(t *testing.T) {
	data := []byte(`{"Category":1,"flag":"test-flag","message":"dGVzdC1tZXNzYWdl"}`)
	envelope, err := UnmarshalEnvelope(data)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if envelope.Category != 1 || envelope.Flag != "test-flag" || string(envelope.Message) != "test-message" {
		t.Errorf("Unexpected envelope content: %+v", envelope)
	}
}

func TestUnmarshalEnvelopeInvalidJSON(t *testing.T) {
	data := []byte(`{"Category":1,"flag":"test-flag","message":`)
	_, err := UnmarshalEnvelope(data)

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
}

func TestUnmarshalEnvelopeEmptyData(t *testing.T) {
	data := []byte(``)
	_, err := UnmarshalEnvelope(data)

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
}

func TestUnmarshalEnvelopeExtraFields(t *testing.T) {
	data := []byte(`{"Category":1,"flag":"test-flag","message":"dGVzdC1tZXNzYWdl","extra":"field"}`)
	envelope, err := UnmarshalEnvelope(data)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if envelope.Category != 1 || envelope.Flag != "test-flag" || string(envelope.Message) != "test-message" {
		t.Errorf("Unexpected envelope content: %+v", envelope)
	}
}
