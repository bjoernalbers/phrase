package passphrase

import (
	"regexp"
	"strings"
	"testing"
)

func TestWordlists(t *testing.T) {
	const wordlistSize = 7776
	validWord := regexp.MustCompile(`\A[a-z-]{3,9}\z`)
	for language, wordlist := range Wordlists {
		t.Run(language, func(t *testing.T) {
			if got := len(wordlist); got != wordlistSize {
				t.Fatalf("wordlist size = %d, want: %d\n", got, wordlistSize)
			}
			for _, word := range wordlist {
				if !validWord.MatchString(word) {
					t.Errorf("invalid word: %q\n", word)
				}
			}
		})
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
	for _, tt := range tests {
		got, err := ReadFile(tt.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ReadFile(%v) error = %v, wantErr: %v", tt.in, err, tt.wantErr)
		}
		if !equal(got, tt.want) {
			t.Errorf("ReadFile(%v) = %v, want: %v", tt.in, got, tt.want)
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
		{
			"ignore duplicate words",
			"11111\tgopher\n11112\tgopher\n",
			[]string{"gopher"},
			false,
		},
		{
			"sort wordlist",
			"11111\tghi\n11112\tabc\n11113\tdef\n",
			[]string{"abc", "def", "ghi"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.in)
			got, _ := read(r)
			if !equal(got, tt.want) {
				t.Errorf("read() = %#v, want: %#v\n", got, tt.want)
			}
		})
	}
}

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
