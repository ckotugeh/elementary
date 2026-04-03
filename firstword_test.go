package main

import (
	"testing"
)

func TestFirstWord(t *testing.T) {
	tests := map[string]string{
		"   hello world": "hello",
		"Go is awesome":  "Go",
		"   ":            "",
		"":               "",
		"singleword":     "singleword",
		" leading space": "leading",
	}

	for input, expected := range tests {
		result := FirstWord(input)
		if result != expected {
			t.Errorf("FirstWord(%q) = %q; want %q", input, result, expected)
		}
	}
}
