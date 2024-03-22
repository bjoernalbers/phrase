package passphrase

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

const WordlistSize = 7776

var validWord = regexp.MustCompile(`\A[a-z]{3,9}\z`)

func equal(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func TestWordlists(t *testing.T) {
	for language, wordlist := range Wordlists {
		if len(wordlist) != WordlistSize {
			t.Fatalf("wordlist[%q] contains %d words, expected %d\n", language, len(wordlist), WordlistSize)
		}
		for _, word := range wordlist {
			if err := ValidateWord(word); err != nil {
				t.Errorf("wordlist[%q] contains invalid word %q\n", language, word)
			}
		}
	}
}

func TestReadFile(t *testing.T) {
	tests := []struct {
		in      string
		want    []string
		wantErr bool
	}{
		{
			"does-not-exist",
			[]string{},
			true,
		},
		{
			"../testdata/gopher.txt",
			[]string{"gopher"},
			false,
		},
	}
	for _, test := range tests {
		got, err := ReadFile(test.in)
		if (err != nil) != test.wantErr {
			t.Errorf("ReadFile(%v) error = %v, wantErr: %v", test.in, err, test.wantErr)
		}
		if !equal(got, test.want) {
			t.Errorf("ReadFile(%v) = %v, want: %v", test.in, got, test.want)
		}
	}
}

func TestRead(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    []string
		wantErr bool
	}{
		{
			"handle empty input",
			"",
			[]string{},
			false,
		},
		{
			"read valid line",
			"11111\tgopher\n",
			[]string{"gopher"},
			false,
		},
		{
			"ignore comment",
			"#11111\tgopher\n",
			[]string{},
			false,
		},
		{
			"ignore empty line",
			"\n",
			[]string{},
			false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := strings.NewReader(test.in)
			got, _ := read(r)
			if !equal(got, test.want) {
				t.Errorf("read() = %#v, want: %#v\n", got, test.want)
			}
		})
	}
}

func ValidateWord(word string) error {
	if !validWord.MatchString(word) {
		return fmt.Errorf("invalid word: %#v", word)
	}
	return nil
}

func TestValidateWord(t *testing.T) {
	tests := []struct {
		name    string
		word    string
		wantErr bool
	}{
		{"two letters", "go", true},
		{"ten letters", "gophergoph", true},
		{"contains space", "go pher", true},
		{"contains umlaut", "göpher", true},
		{"contains hyphen", "go-pher", true},
		{"contains dot", "go.pher", true},
		{"contains digit", "goph3r", true},
		{"all letters uppercase", "GOPHER", true},
		{"first letter capitalized", "Gopher", true},
		{"last letter capitalized", "gopheR", true},
		{"all letters lowercase", "gopher", false},
		{"three letters", "gop", false},
		{"nine letters", "gophergop", false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := ValidateWord(test.word); (err != nil) != test.wantErr {
				t.Fatalf("validateWord(%#v) = %v, wantErr: %#v\n", test.word, err, test.wantErr)
			}
		})
	}
}
