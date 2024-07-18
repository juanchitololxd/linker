package domain

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestURLMappingJSONSerialization(t *testing.T) {
	original := URLMapping{
		OriginalURL: "http://example.com",
		ShortURL:    "http://1.unli.ink/s/abc123",
	}

	// Serialize to JSON
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("failed to marshal URLMapping: %v", err)
	}

	expectedJSON := `{"original_url":"http://example.com","short_url":"http://1.unli.ink/s/abc123"}`
	if string(data) != expectedJSON {
		t.Errorf("unexpected JSON serialization: got %s, want %s", data, expectedJSON)
	}

	// Deserialize from JSON
	var deserialized URLMapping
	if err := json.Unmarshal(data, &deserialized); err != nil {
		t.Fatalf("failed to unmarshal URLMapping: %v", err)
	}

	if !reflect.DeepEqual(original, deserialized) {
		t.Errorf("unexpected deserialization: got %+v, want %+v", deserialized, original)
	}
}

func TestURLMappingJSONDeserialization(t *testing.T) {
	jsonData := `{"original_url":"http://example.com","short_url":"http://1.unli.ink/s/abc123"}`

	var urlMapping URLMapping
	if err := json.Unmarshal([]byte(jsonData), &urlMapping); err != nil {
		t.Fatalf("failed to unmarshal URLMapping: %v", err)
	}

	expected := URLMapping{
		OriginalURL: "http://example.com",
		ShortURL:    "http://1.unli.ink/s/abc123",
	}

	if !reflect.DeepEqual(urlMapping, expected) {
		t.Errorf("unexpected deserialization: got %+v, want %+v", urlMapping, expected)
	}
}
