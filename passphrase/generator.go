package passphrase

import (
	"crypto/rand"
	"fmt"
	"math"
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
	var passphrase []string
	for i := 0; i < g.Words; i++ {
		r, err := randInt(len(g.Wordlist))
		if err != nil {
			return "", err
		}
		passphrase = append(passphrase, g.Wordlist[r])
	}
	if g.Capitalize {
		for i := range passphrase {
			passphrase[i] = strings.Title(passphrase[i])
		}
	}
	if g.Digits > 0 {
		var n int
		for i := 0; i < g.Digits; i++ {
			r, err := randInt(10)
			if err != nil {
				return "", err
			}
			n = n + int(math.Pow10(i))*r
		}
		passphrase = append(passphrase, strconv.Itoa(n))
	}
	return strings.Join(passphrase, g.Separator), nil
}

// randInt returns a random integer in the range [0, max).
// Errors from the underlying crypto function get passed through.
func randInt(max int) (int, error) {
	random, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, fmt.Errorf("random number generator: %v", err)
	}
	return int(random.Int64()), nil
}
