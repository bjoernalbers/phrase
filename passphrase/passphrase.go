// Package passphrase generates easy-to-remember passwords from random words.
package passphrase

import (
	"bufio"
	"io"
	"os"
	"regexp"
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
func read(reader io.Reader) (wordlist []string, err error) {
	buffer := make(map[string]bool)
	validLine := regexp.MustCompile(`\A[1-6]{5}\t(.+)\z`)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		match := validLine.FindStringSubmatch(scanner.Text())
		if len(match) == 0 {
			continue
		}
		buffer[match[1]] = true
	}
	if err := scanner.Err(); err != nil {
		return wordlist, scanner.Err()
	}
	for word := range buffer {
		wordlist = append(wordlist, word)
	}
	return wordlist, nil
}
