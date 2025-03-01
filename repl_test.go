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
		}, {
			input:    "my name jeff ",
			expected: []string{"my", "name", "jeff"},
		}, {
			input:    " you belong in     a museum",
			expected: []string{"you", "belong", "in", "a", "museum"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) returned %d words, expected %d words",
				c.input, len(actual), len(c.expected))
		}

		for i := range actual {
			if i < len(c.expected) && actual[i] != c.expected[i] {
				t.Errorf("cleanInput(%q) word %d: got %q, expected %q",
					c.input, i, actual[i], c.expected[i])
			}
		}
	}
}
