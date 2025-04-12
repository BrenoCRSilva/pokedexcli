package repl

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "Hello, World",
			expected: []string{"hello,", "world"},
		},
		{
			input:    "make love not war",
			expected: []string{"make", "love", "not", "war"},
		},
		{
			input:    " you BROKE  me First",
			expected: []string{"you", "broke", "me", "first"},
		},
	}
	for _, c := range cases {
		current := cleanInput(c.input)
		if len(current) != len(c.expected) {
			t.Errorf("Input length does not match")
		}
		for i, word := range current {
			if word != c.expected[i] {
				t.Errorf("%s does not match %s", word, c.expected[i])
			}
		}

	}
}
