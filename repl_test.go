package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "this space this	tab",
			expected: []string{"this", "space", "this", "tab"},
		},
		{
			input:    "under_score and end     ",
			expected: []string{"under_score", "and", "end"},
		},
		{
			input:    "CAPS CHECK",
			expected: []string{"caps", "check"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf(
				"test failed: actual slice length [%d] and expected slice length [%d] do not match", len(actual), len(c.expected))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("test failed: actual word [%s] does not match expected word [%s]", word, expectedWord)
			}
		}
	}
}
