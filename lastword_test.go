package main

import (
	"testing"
)

func TestLastWord(t *testing.T) {
	tests := map[string]string{
		"hello world":         "world",
		"Go is awesome":       "awesome",
		"   multiple spaces ": "spaces",
		"":                    "",
		"singleword":          "singleword",
		" trailing space ":    "space",
	}

	for input, expected := range tests {
		result := LastWord(input)
		if result != expected {
			t.Errorf("LastWord(%q) = %q; want %q", input, result, expected)
		}
	}
}
