package main

import (
	"testing"
)

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

func TestPick(t *testing.T) {
	words := []string{"gopher"}
	tests := []struct {
		in      int
		want    []string
		wantErr bool
	}{
		{
			-1,
			[]string{},
			true,
		},
		{
			0,
			[]string{},
			true,
		},
		{
			1,
			[]string{"gopher"},
			false,
		},
		{
			2,
			[]string{"gopher", "gopher"},
			false,
		},
	}
	for _, test := range tests {
		got, err := pick(words, test.in)
		if (err != nil) != test.wantErr {
			t.Errorf("pick(words, %v) error = %v, wantErr: %v", test.in, err, test.wantErr)
		}
		if !equal(got, test.want) {
			t.Errorf("pick(%v, %v) = %v, want: %v", words, test.in, got, test.want)
		}
	}
	words = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	randomWords1, _ := pick(words, len(words))
	randomWords2, _ := pick(words, len(words))
	if equal(randomWords1, randomWords2) {
		t.Errorf("pick() returned the same result twice")
	}
}
