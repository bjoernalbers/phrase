package passphrase

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// Generator generates random passphrases.
type Generator struct {
	Wordlist   []string
	Words      int
	Separator  string
	Capitalize bool
	Digits     int
}

// Phrase returns a random passphrase or an error, if the randomization fails.
func (g *Generator) Phrase() (string, error) {
	passphrase, err := randomWords(g.Wordlist, g.Words)
	if err != nil {
		return "", err
	}
	if g.Capitalize {
		for i := range passphrase {
			passphrase[i] = strings.Title(passphrase[i])
		}
	}
	if g.Digits > 0 {
		var number string
		for i := 0; i < g.Digits; i++ {
			r, err := randomInt(10)
			if err != nil {
				return "", err
			}
			// Numbers with multiple digits should not begin with
			// zeroes.
			if g.Digits > 1 && i == 0 && r == 0 {
				r = 1
			}
			number += strconv.Itoa(r)
		}
		passphrase = append(passphrase, number)
	}
	return strings.Join(passphrase, g.Separator), nil
}

// randomInt returns a random integer in the range [0, max), where max must be
// greater than zero.
func randomInt(max int) (int, error) {
	if max <= 0 {
		return 0, fmt.Errorf("randomInt: max must be greater than zero")
	}
	r, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, fmt.Errorf("randomInt: %v", err)
	}
	return int(r.Int64()), nil
}

// randomWords returns n random words from wordlist.
func randomWords(wordlist []string, n int) (words []string, err error) {
	length := len(wordlist)
	if length == 0 {
		return words, nil
	}
	for i := 0; i < n; i++ {
		r, err := randomInt(length)
		if err != nil {
			return words, err
		}
		words = append(words, wordlist[r])
	}
	return words, nil
}
