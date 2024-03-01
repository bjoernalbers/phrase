// phrase - the passphrase generator

//go:generate go run generate.go

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/bjoernalbers/phrase/passphrase"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix(fmt.Sprintf("%s: ", filepath.Base(os.Args[0])))
}

func main() {
	g := passphrase.Generator{}
	flag.IntVar(&g.Words, "w", 4, "Number of words per passphrase.")
	flag.StringVar(&g.Separator, "s", " ", "Separator between words.")
	flag.BoolVar(&g.Capitalize, "C", false, "Capitalize all words")
	flag.IntVar(&g.Digits, "d", 0, "Digits per passphrase.")
	filename := flag.String("f", "", "Diceware wordlist file.")
	language := flag.String("l", "de", "Language of wordlist.")
	flag.Parse()
	var err error
	if *filename != "" {
		g.Wordlist, err = passphrase.ReadFile(*filename)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		var ok bool
		g.Wordlist, ok = passphrase.Wordlists[*language]
		if !ok {
			log.Fatalf("no such language: %q", *language)
		}
	}
	passphrase, err := g.Phrase()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(passphrase)
}
