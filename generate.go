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
	directory, filename := filepath.Split(inputPath)
	language := strings.TrimSuffix(filename, filepath.Ext(filename))
	outputPath := filepath.Join(directory, language+".go")
	wordlist, err := wordlists.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}
	err = templ.Execute(file, Wordlist{language, wordlist})
	if err != nil {
		log.Fatal(err)
	}
}
