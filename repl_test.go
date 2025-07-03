package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " fuck you ",
			expected: []string{"fuck", "you"},
		},
		{
			input:    " Barrack Obama loves his manwife",
			expected: []string{"barrack", "obama", "loves", "his", "manwife"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("ERR - Expected length: %v - Actual length : %v", len(c.expected), len(actual))
		}
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("ERR - Expected: %v - Actual: %v", expectedWord, word)
			}
		}
	}

}
