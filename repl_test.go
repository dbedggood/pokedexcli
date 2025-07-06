package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {

	type testCase struct {
		input    string
		expected []string
	}

	testCases := []testCase{
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  \t\n mixed   whitespace \t\n ",
			expected: []string{"mixed", "whitespace"},
		},
	}

	for _, tc := range testCases {
		actual := cleanInput(tc.input)

		if len(actual) != len(tc.expected) {
			t.Errorf("cleanInput(%q) returned %d elements; expected %d", tc.input, len(actual), len(tc.expected))
		}

		for i := range actual {
			if actual[i] != tc.expected[i] {
				t.Errorf("cleanInput(%q) = %v; expected %v", tc.input, actual, tc.expected)
			}
		}
	}

}
