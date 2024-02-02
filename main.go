// phrase - the passphrase generator
package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/bjoernalbers/phrase/wordlists"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix(fmt.Sprintf("%s: ", filepath.Base(os.Args[0])))
}

func main() {
	filename := flag.String("f", "", "Diceware wordlist file.")
	flag.Parse()
	var wordlist []string
	var err error
	if *filename != "" {
		wordlist, err = wordlists.ReadFile(*filename)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// TODO: Replace fake data with real wordlist.
		wordlist = []string{"correct", "horse", "battery", "staple"}
	}
	passphrase, err := pick(wordlist, 4)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strings.Join(passphrase, " "))
}

// pick returns a slice of n random words from words.
func pick(words []string, n int) (randomWords []string, err error) {
	if n <= 0 {
		return nil, fmt.Errorf("Number of words to pick must be greater than 0")
	}
	for i := 0; i < n; i++ {
		randomWords = append(randomWords, words[rand.Intn(len(words))])
	}
	return randomWords, nil
}
