// phrase - the passphrase generator
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
)

var words = []string{
	"correct",
	"horse",
	"battery",
	"staple",
}

func main() {
	flag.Parse()
	randomWords, _ := pick(words, 4)
	fmt.Println(strings.Join(randomWords, " "))
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
