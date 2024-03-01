// phrase - the passphrase generator

//go:generate go run generate.go

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/bjoernalbers/phrase/wordlists"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix(fmt.Sprintf("%s: ", filepath.Base(os.Args[0])))
}

func main() {
	g := Generator{}
	flag.IntVar(&g.Words, "w", 4, "Number of words per passphrase.")
	flag.StringVar(&g.Separator, "s", " ", "Separator between words.")
	flag.BoolVar(&g.Capitalize, "C", false, "Capitalize all words")
	flag.IntVar(&g.Digits, "d", 0, "Digits per passphrase.")
	filename := flag.String("f", "", "Diceware wordlist file.")
	language := flag.String("l", "de", "Language of wordlist.")
	flag.Parse()
	var err error
	if *filename != "" {
		g.Wordlist, err = wordlists.ReadFile(*filename)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		g.Wordlist, err = wordlists.Get(*language)
		if err != nil {
			log.Fatal(err)
		}
	}
	passphrase, err := g.Phrase()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(passphrase)
}
