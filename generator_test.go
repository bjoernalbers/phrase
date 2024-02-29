package main

import (
	"regexp"
	"testing"
)

func TestGenerator(t *testing.T) {
	tests := []struct {
		name   string
		in     Generator
		want   string
		regexp bool
	}{
		{
			"default generator",
			Generator{},
			"",
			false,
		},
		{
			"change number of words",
			Generator{Wordlist: []string{"gopher"}, Words: 1},
			"gopher",
			false,
		},
		{
			"change separator",
			Generator{Wordlist: []string{"gopher"}, Words: 2, Separator: " "},
			"gopher gopher",
			false,
		},
		{
			"capitalize words",
			Generator{Wordlist: []string{"gopher"}, Words: 1, Capitalize: true},
			"Gopher",
			false,
		},
		{
			"change number of digits",
			Generator{Digits: 10},
			`^[0-9]{10}$`,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.in.Phrase()
			if err != nil {
				t.Fatalf("Generator.Phrase() error = %v", err)
			}
			if tt.regexp {
				if !regexp.MustCompile(tt.want).MatchString(got) {
					t.Errorf("Generator.Phrase() = %q, does not match: %q", got, tt.want)
				}
			} else {
				if got != tt.want {
					t.Errorf("Generator.Phrase() = %q, want: %q", got, tt.want)
				}
			}
		})
	}
}
