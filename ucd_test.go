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
		// Unicode 16
		"🫩": "FACE WITH BAGS UNDER EYES",
		// Unicode 17
		"🫍": "ORCA",
		// Unicode 17 Unihan
		"方言字": "A SQUARE, RECTANGLE; A REGION; LOCAL",
	}

	for raw, expected := range tests {

		name := Name(raw)

		if name.String() != expected {
			t.Fatalf("Failed to derive name for '%s'. Expected '%s' but got '%s'", raw, expected, name.String())
		}
	}

}
