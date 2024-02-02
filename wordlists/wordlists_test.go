package wordlists

import "testing"

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
