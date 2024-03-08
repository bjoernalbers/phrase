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
	flag.StringVar(&g.Language, "l", "de", "Language of wordlist.")
	filename := flag.String("f", "", "Diceware wordlist file.")
	passphrases := flag.Int("p", 1, "Number of passphrases")
	flag.Parse()
	var err error
	if *filename != "" {
		g.Wordlist, err = passphrase.ReadFile(*filename)
		if err != nil {
			log.Fatal(err)
		}
	}
	for i := 0; i < *passphrases; i++ {
		passphrase, err := g.Phrase()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(passphrase)
	}
}
