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
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    " Hola     a    todos",
			expected: []string{"hola", "a", "todos"},
		},
		{
			input:    "HelLO       WOrlD",
			expected: []string{"hello", "world"},
		},
	}
	for _, c := range cases {
		// Check the lenght of the actual slice against the expected slice
		// if the don't match, use fmt.Errorf to print an error message
		// and fail the test
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("the lenghts do not match '%v' vs '%v'", actual, c.expected)
			continue
		}

		// Check each word in the slice
		// if they don't match, use fmt.Errorf to print error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}

		}
	}
}
