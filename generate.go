// generate.go generates Go wordlists from diceware wordlists.
//
// Usage: go run generate.go

//go:build ignore

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

const (
	dir          = "wordlists"
	inputSuffix  = ".txt"
	outputSuffix = ".go"
	templateFile = "wordlists/wordlist.go.tmpl"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix(fmt.Sprintf("%s: ", filepath.Base(os.Args[0])))
}

func main() {
	files, err := find(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		err = generate(file)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// find returns the names of all input files inside dir.
func find(dir string) (files []string, err error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return files, err
	}
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == inputSuffix {
			files = append(files, filepath.Join(dir, entry.Name()))
		}
	}
	return files, err
}

// generate generates a Go wordlist from a diceware wordlist.
func generate(input string) error {
	dirname, filename := filepath.Split(input)
	language := strings.TrimSuffix(filename, filepath.Ext(filename))
	output := filepath.Join(dirname, language+outputSuffix)
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
	templ, err := template.New(filepath.Base(templateFile)).ParseFiles(templateFile)
	if err != nil {
		return err
	}
	err = templ.Execute(file, Wordlist{language, wordlist})
	if err != nil {
		return err
	}
	fmt.Println(output)
	return nil
}
