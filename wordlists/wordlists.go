// Package wordlists provides lists of words to generate passphrases from.
package wordlists

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

const ValidWordRegexp = `\A[a-z]{3,9}\z`

// Lists contains all wordlists grouped by two-letter language code.
var Lists map[string][]string

func init() {
	Lists = make(map[string][]string)
}

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
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		_, word, _ := strings.Cut(scanner.Text(), "\t")
		wordlist = append(wordlist, word)
	}
	return wordlist, err
}

func ValidateWord(word string) error {
	if matched, _ := regexp.MatchString(ValidWordRegexp, word); !matched {
		return fmt.Errorf("invalid word: %#v", word)
	}
	return nil
}
