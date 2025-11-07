package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  ",
			expected: []string{},
		},
		{
			input: "  hello  ",
			expected: []string{"hello"},
		},
		{
			input: "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Hello WORLD",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("%v is not the same length as expected output",actual)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("%v does not equal %v", word, expectedWord)
			}
		}
	}
	
}
	
