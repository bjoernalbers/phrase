// Package passphrase provides tools to generate passphrases,
// that is passwords made from random words.
package passphrase

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"
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
	var dicewarePrefix = regexp.MustCompile(`\A[1-6]{5}\t`)
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
