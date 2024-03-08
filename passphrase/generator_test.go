package passphrase

import (
	"regexp"
	"testing"
)

func TestGenerator(t *testing.T) {
	tests := []struct {
		name string
		in   Generator
		want string
	}{
		{
			"default generator",
			Generator{},
			"",
		},
		{
			"change number of words",
			Generator{Wordlist: []string{"gopher"}, Words: 1},
			"gopher",
		},
		{
			"change separator",
			Generator{Wordlist: []string{"gopher"}, Words: 2, Separator: " "},
			"gopher gopher",
		},
		{
			"capitalize words",
			Generator{Wordlist: []string{"gopher"}, Words: 1, Capitalize: true},
			"Gopher",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.in.Phrase()
			if err != nil {
				t.Fatalf("Generator.Phrase() error = %v", err)
			}
			if got != tt.want {
				t.Errorf("Generator.Phrase() = %q, want: %q", got, tt.want)
			}
		})
	}
	tests = []struct {
		name string
		in   Generator
		want string
	}{
		{
			"one digit",
			Generator{Digits: 1},
			`^[0-9]$`,
		},
		{
			"multiple digits",
			Generator{Digits: 10},
			`^[1-9][0-9]{9}$`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 100; i++ {
				got, err := tt.in.Phrase()
				if err != nil {
					t.Fatalf("Generator.Phrase() error = %v", err)
				}
				if !regexp.MustCompile(tt.want).MatchString(got) {
					t.Fatalf("Generator.Phrase() = %q, does not match: %q", got, tt.want)
				}
			}
		})
	}
}

func TestRandomInt(t *testing.T) {
	tests := []struct {
		in      int
		wantErr bool
	}{
		{-1, true},
		{0, true},
		{1, false},
	}
	for _, tt := range tests {
		if _, err := randomInt(tt.in); (err != nil) != tt.wantErr {
			t.Errorf("randomInt(%d) error = %v, wantErr: %v", tt.in, err, tt.wantErr)
		}
	}
}
