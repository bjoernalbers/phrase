package passphrase

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

// Generator generates random passphrases by the given settings.
type Generator struct {
	// Wordlist must be a list of words from which the passphrases are
	// generated. It can be set directly, i.e. to the output of ReadFile().
	// Or it can be set indirectly via the Language field (see below).
	Wordlist []string

	// Language allows to set Wordlist to a build-in wordlist, which must
	// be present in Wordlists.
	// The Language setting gets ignored when Wordlist has already been set
	// to a non-empty value.
	Language string

	// Words determines the number of words per passphrase.
	Words int

	// Separator is the string between the words that joins all words
	// together. It is usually a single character or an empty string.
	Separator string

	// Capitalize converts the first letter of each word to an uppercase
	// letter when set to true.
	Capitalize bool

	// Digits determines the number of digits per passphrase.
	Digits int
}

// Phrase returns a random passphrase.
func (g *Generator) Phrase() (string, error) {
	if g.Language != "" && len(g.Wordlist) == 0 {
		wordlist, ok := Wordlists[g.Language]
		if !ok {
			return "", fmt.Errorf("no such language: %s", g.Language)
		}
		g.Wordlist = wordlist
	}
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
		number, err := randomNumber(g.Digits)
		if err != nil {
			return "", err
		}
		passphrase = append(passphrase, strconv.Itoa(number))
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

// randomNumber returns a random number with n digits.
func randomNumber(n int) (int, error) {
	var min, max int
	if n == 1 {
		min = 0
	} else {
		min = int(math.Pow10(n - 1))
	}
	max = int(math.Pow10(n)) - min
	r, err := randomInt(max)
	return r + min, err
}
