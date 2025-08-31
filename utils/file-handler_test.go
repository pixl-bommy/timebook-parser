package utils

import (
	"reflect"
	"testing"
)

func TestRetrieveLinesFromContent(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []string
	}{
		{
			name:     "Empty input",
			input:    []byte(""),
			expected: []string{},
		},
		{
			name:     "Single line, no newline",
			input:    []byte("hello world"),
			expected: []string{"hello world"},
		},
		{
			name:     "Single line with newline",
			input:    []byte("hello world\n"),
			expected: []string{"hello world"},
		},
		{
			name:     "Multiple lines with newlines",
			input:    []byte("line1\nline2\nline3\n"),
			expected: []string{"line1", "line2", "line3"},
		},
		{
			name:     "Multiple lines, last line without newline",
			input:    []byte("line1\nline2\nline3"),
			expected: []string{"line1", "line2", "line3"},
		},
		{
			name:     "Lines with carriage returns",
			input:    []byte("line1\r\nline2\r\nline3\r\n"),
			expected: []string{"line1", "line2", "line3"},
		},
		{
			name:     "Consecutive newlines (empty lines)",
			input:    []byte("line1\n\nline3\n"),
			expected: []string{"line1", "", "line3"},
		},
		{
			name:     "Carriage return at end",
			input:    []byte("line1\r"),
			expected: []string{"line1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RetrieveLinesFromContent(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("RetrieveLinesFromContent(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
