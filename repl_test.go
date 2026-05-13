package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "only spaces",
			input:    "     ",
			expected: []string{},
		},
		{
			name:     "single word",
			input:    "Pikachu",
			expected: []string{"pikachu"},
		},
		{
			name:     "multiple spaces between words",
			input:    "hello     world",
			expected: []string{"hello", "world"},
		},
		{
			name:     "leading and trailing spaces",
			input:    "   bulbasaur charmander   ",
			expected: []string{"bulbasaur", "charmander"},
		},
		{
			name:     "tabs and newlines",
			input:    "\thello\nworld\tpikachu",
			expected: []string{"hello", "world", "pikachu"},
		},
		{
			name:     "mixed uppercase lowercase",
			input:    "HeLLo WoRLD",
			expected: []string{"hello", "world"},
		},
		{
			name:     "special characters",
			input:    "hello, world!",
			expected: []string{"hello,", "world!"},
		},
		{
			name:     "numbers",
			input:    "poke123 test456",
			expected: []string{"poke123", "test456"},
		},
		{
			name:     "unicode accents",
			input:    "Pokémon Café Açúcar",
			expected: []string{"pokémon", "café", "açúcar"},
		},
		{
			name:     "emoji",
			input:    "pikachu ⚡ charmander 🔥",
			expected: []string{"pikachu", "⚡", "charmander", "🔥"},
		},
		{
			name:     "multiple blank lines",
			input:    "\n\nhello\n\nworld\n",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths mismatch: '%v' x '%v'", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}
