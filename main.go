// phrase - the passphrase generator

//go:generate go run generate.go

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bjoernalbers/phrase/passphrase"
)

var (
	version  = "unset" // version gets set via build flags
	homepage = "https://github.com/bjoernalbers/phrase"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix(fmt.Sprintf("%s: ", filepath.Base(os.Args[0])))
	flag.Usage = usage
}

func main() {
	g := passphrase.Generator{}
	flag.IntVar(&g.Words, "w", 4, "Words per passphrase")
	flag.StringVar(&g.Separator, "s", " ", "Separator between words")
	flag.BoolVar(&g.Capitalize, "C", false, "Capitalize all words")
	flag.IntVar(&g.Digits, "d", 0, "Digits per passphrase")
	flag.StringVar(&g.Language, "l", "en", "Language of passphrase: "+strings.Join(languages(), ", "))
	filename := flag.String("f", "", "Diceware wordlist file")
	passphrases := flag.Int("p", 1, "Passphrases")
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

func usage() {
	header := fmt.Sprintf(`phrase - the passphrase generator (version %s)

Generate easy-to-remember passwords from random words.

Usage:
`, version)
	footer := fmt.Sprintf(`
Homepage: %s
`, homepage)
	fmt.Fprintf(flag.CommandLine.Output(), header)
	flag.PrintDefaults()
	fmt.Fprintf(flag.CommandLine.Output(), footer)
}

// languages returns the names of all build-in wordlists
func languages() []string {
	list := []string{}
	for language := range passphrase.Wordlists {
		list = append(list, language)
	}
	return list
}
