package main

import "testing"

func TestDetectWordsFromCamelCase(t *testing.T) {
	testCases := []struct {
		in  string
		out int
	}{
		{"thisIsASimpleSentense", 5},
	}

	for _, tc := range testCases {
		got := detectWordsFromCamelCase(tc.in)
		if got != tc.out {
			t.Errorf("Expected %v, but got %d", tc.out, got)
		}
	}
}
