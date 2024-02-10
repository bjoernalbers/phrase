// phrase - the passphrase generator

//go:generate go run generate.go

package main

import (
	"flag"
	"fmt"
	"log"
	"math"
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
	language := flag.String("l", "de", "Language of wordlist.")
	words := flag.Int("w", 4, "Number of words per passphrase.")
	separator := flag.String("s", " ", "Separator between words.")
	digits := flag.Int("d", 0, "Digits per passphrase.")
	flag.Parse()
	var wordlist []string
	var err error
	if *filename != "" {
		wordlist, err = wordlists.ReadFile(*filename)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		wordlist, err = wordlists.Get(*language)
		if err != nil {
			log.Fatal(err)
		}
	}
	passphrase, err := pick(wordlist, *words)
	if err != nil {
		log.Fatal(err)
	}
	if *digits > 0 {
		passphrase = append(passphrase, fmt.Sprintf("%d", rand.Intn(int(math.Pow10(*digits)))))
	}
	fmt.Println(strings.Join(passphrase, *separator))
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
