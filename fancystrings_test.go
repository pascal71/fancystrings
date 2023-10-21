// fancystrings/fancystrings_test.go

package fancystrings

import (
	"testing"
)

func TestInsertSpaces(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty String",
			input:    "",
			expected: "",
		},
		{
			name:     "Single Character",
			input:    "H",
			expected: "H",
		},
		{
			name:     "Multiple Characters",
			input:    "Hello",
			expected: "H e l l o",
		},
		{
			name:     "With Numbers and Symbols",
			input:    "Hi!123",
			expected: "H i ! 1 2 3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InsertSpaces(tt.input)
			if got != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, got)
			}
		})
	}
}
