// phrase - the passphrase generator
package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("phrase: ")
}

func main() {
	list := flag.String("l", "", "List of words from which the passphrases are generated. The list must be a path to a diceware wordlist.")
	flag.Parse()
	words, err := readList(*list)
	if err != nil {
		log.Fatal(err)
	}
	randomWords, err := pick(words, 4)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strings.Join(randomWords, " "))
}

// readList reads and returns words from list.
func readList(list string) ([]string, error) {
	if list == "" {
		return []string{"correct", "horse", "battery", "staple"}, nil
	}
	_, err := os.Open(list)
	if err != nil {
		return []string{}, err
	}
	return []string{"hello"}, nil
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
