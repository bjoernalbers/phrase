// phrase - the passphrase generator
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/bjoernalbers/phrase/wordlists"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("phrase: ")
}

func main() {
	filename := flag.String("f", "", "Diceware wordlist file.")
	flag.Parse()
	words, err := readList(*filename)
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
	words, ok := wordlists.Lists[list]
	if ok {
		return words, nil
	}
	f, err := os.Open(list)
	if err != nil {
		return []string{}, err
	}
	defer f.Close()
	words, err = read(f)
	return words, nil
}

// read reads and returns words of r
func read(r io.Reader) (words []string, err error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		_, word, _ := strings.Cut(scanner.Text(), "\t")
		words = append(words, word)
	}
	return words, err
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
