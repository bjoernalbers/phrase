//go:build ignore

package main

import (
	"os/exec"
	"regexp"
	"strings"
	"testing"
)

const binary = "./phrase"

// run executes the binary under test with the given arguments and returns its
// combined output and error.
func run(args ...string) ([]byte, error) {
	return exec.Command(binary, args...).CombinedOutput()
}

// equal returns true if both slices are equal, otherwise false.
func equal(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestEqual(t *testing.T) {
	tests := []struct {
		a, b []byte
		want bool
	}{
		{
			[]byte{},
			[]byte{},
			true,
		},
		{
			[]byte{1, 2, 3},
			[]byte{1, 2, 3},
			true,
		},
		{
			[]byte{1, 2, 3},
			[]byte{1, 2, 3, 4},
			false,
		},
		{
			[]byte{1, 2, 3},
			[]byte{1, 3, 2},
			false,
		},
	}
	for _, tt := range tests {
		if got := equal(tt.a, tt.b); got != tt.want {
			t.Fatalf("equal(%v, %v) = %v, want: %v\n", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestRandomness(t *testing.T) {
	first, err := run()
	if err != nil {
		t.Fatalf("Command returned an error: %v\n%s", err, first)
	}
	second, err := run()
	if err != nil {
		t.Fatalf("Command returned an error: %v\n%s", err, second)
	}
	if equal(first, second) {
		t.Fatalf("Command returned non-random output")
	}
}

func TestIntegration(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    string
		wantErr bool
	}{
		{"default output", nil, `^[a-z]{3,9} [a-z]{3,9} [a-z]{3,9} [a-z]{3,9}\n$`, false},
		{"display help", []string{"-h"}, `Usage`, false},
		{"display version", []string{"-h"}, `(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?`, false},
		{"change wordlist", []string{"-f", "testdata/gopher.txt"}, `^gopher gopher gopher gopher\n$`, false},
		{"change number of words", []string{"-w", "3", "-f", "testdata/gopher.txt"}, `^gopher gopher gopher\n$`, false},
		{"change number of passphrases", []string{"-p", "3", "-w", "1", "-f", "testdata/gopher.txt"}, `^gopher\ngopher\ngopher\n$`, false},
		{"change separator", []string{"-s", "-", "-f", "testdata/gopher.txt"}, `^gopher-gopher-gopher-gopher\n$`, false},
		{"add digits", []string{"-d", "3"}, `\W[0-9]{3}\W`, false},
		{"capitalize words", []string{"-C", "-f", "testdata/gopher.txt"}, `^Gopher Gopher Gopher Gopher\n$`, false},
		{"missing wordlist", []string{"-f", "this-file-does-not-exist"}, `this-file-does-not-exist`, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := run(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Unexpected exit of \"%s %s\":\n%s", binary, strings.Join(tt.args, " "), got)
			}
			if !regexp.MustCompile(tt.want).Match(got) {
				t.Fatalf("Output %q did not match %q.\n", got, tt.want)
			}
		})
	}
}
