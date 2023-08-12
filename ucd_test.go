package ucd

import (
	"testing"
)

func TestName(t *testing.T) {

	tests := map[string]string{
		// Unicode 14.0
		"🪺": "NEST WITH EGGS",
		// Unicode 15.0
		"🪿": "GOOSE",
	}

	for raw, expected := range tests {

		name := Name(raw)

		if name.String() != expected {
			t.Fatalf("Failed to derive name for '%s'. Expected '%s' but got '%s'", raw, expected, name.String())
		}
	}

}
