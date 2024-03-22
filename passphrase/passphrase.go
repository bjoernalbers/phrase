// Package passphrase generates... *drumroll*... passphrases.
package passphrase

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var (
	validWord      = regexp.MustCompile(`\A[a-z]{3,9}\z`)
	dicewarePrefix = regexp.MustCompile(`\A[1-6]{5}\t`)
)

// Wordlists contains all wordlists grouped by language.
// Each new language file add a wordlist to this map with the corresponding
// two-letter language code as key.
var Wordlists = map[string][]string{}

// ReadFile reads and returns wordlist from filename.
func ReadFile(filename string) (wordlist []string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer f.Close()
	wordlist, err = read(f)
	return wordlist, nil
}

// read reads and returns wordlist from reader.
func read(reader io.Reader) ([]string, error) {
	var wordlist []string
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if !dicewarePrefix.MatchString(scanner.Text()) {
			continue
		}
		_, word, _ := strings.Cut(scanner.Text(), "\t")
		wordlist = append(wordlist, word)
	}
	if err := scanner.Err(); err != nil {
		return wordlist, scanner.Err()
	}
	return wordlist, nil
}

func ValidateWord(word string) error {
	if !validWord.MatchString(word) {
		return fmt.Errorf("invalid word: %#v", word)
	}
	return nil
}
