//go:build ignore

// generate.go generates Go wordlists from diceware wordlists.
//
// Usage: go run generate.go wordlists/foo.txt
//
// The former command would generate "wordlists/foo.go" with a language of "foo".
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/bjoernalbers/phrase/wordlists"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix(fmt.Sprintf("%s: ", filepath.Base(os.Args[0])))
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Please provide a diceware wordlist file!")
	}
	inputPath := os.Args[1]
	err := generate(inputPath)
	if err != nil {
		log.Fatal(err)
	}
}

// generate generates a Go wordlist from a diceware wordlist.
func generate(input string) error {
	dirname, filename := filepath.Split(input)
	language := strings.TrimSuffix(filename, filepath.Ext(filename))
	output := filepath.Join(dirname, language+".go")
	wordlist, err := wordlists.ReadFile(input)
	if err != nil {
		return err
	}
	file, err := os.Create(output)
	if err != nil {
		return err
	}
	defer file.Close()
	type Wordlist struct {
		Language string
		Wordlist []string
	}
	templ, err := template.New("wordlist").Parse(`package wordlists

func init() {
	wordlists["{{.Language}}"] = []string{
{{range .Wordlist}}		"{{.}}",
{{end}}	}
}
`)
	if err != nil {
		return err
	}
	err = templ.Execute(file, Wordlist{language, wordlist})
	if err != nil {
		return err
	}
	return nil
}
