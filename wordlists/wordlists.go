// Package wordlists provides lists of words to generate passphrases from.
package wordlists

import (
	"fmt"
	"regexp"
)

const ValidWordRegexp = `\A[a-z]{3,9}\z`

// Lists contains all wordlists grouped by two-letter language code.
var Lists map[string][]string

func init() {
	Lists = make(map[string][]string)

	// TODO: Replace fake data with real wordlist.
	Lists[""] = []string{"correct", "horse", "battery", "staple"}
}

func ValidateWord(word string) error {
	if matched, _ := regexp.MatchString(ValidWordRegexp, word); !matched {
		return fmt.Errorf("invalid word: %#v", word)
	}
	return nil
}
