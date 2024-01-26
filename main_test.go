package main

import (
	"strings"
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

func TestReadList(t *testing.T) {
	tests := []struct {
		list    string
		want    []string
		wantErr bool
	}{
		{
			"",
			[]string{"correct", "horse", "battery", "staple"},
			false,
		},
		{
			"does-not-exist",
			[]string{},
			true,
		},
		{
			"testdata/gopher.txt",
			[]string{"gopher"},
			false,
		},
	}
	for _, test := range tests {
		got, err := readList(test.list)
		if (err != nil) != test.wantErr {
			t.Errorf("readList(%v) error = %v, wantErr: %v", test.list, err, test.wantErr)
		}
		if !equal(got, test.want) {
			t.Errorf("readList(%v) = %v, want: %v", test.list, got, test.want)
		}
	}
}

func TestRead(t *testing.T) {
	tests := []struct {
		in      string
		want    []string
		wantErr bool
	}{
		{
			"",
			[]string{},
			false,
		},
		{
			"11111\tgopher\n",
			[]string{"gopher"},
			false,
		},
	}
	for _, test := range tests {
		r := strings.NewReader(test.in)
		got, _ := read(r)
		if !equal(got, test.want) {
			t.Errorf("read() = %v, want: %v\n", got, test.want)
		}
	}
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
