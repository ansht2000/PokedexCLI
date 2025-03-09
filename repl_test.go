package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello     world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "helloworld",
			expected: []string{"helloworld"},
		},
		{
			input:    "pokemon charizard BULBAsaur Squirtle",
			expected: []string{"pokemon", "charizard", "bulbasaur", "squirtle"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word: %s, does not match expected word: %s", word, expectedWord)
				return
			}
		}
	}
}
