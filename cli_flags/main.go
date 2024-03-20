package main

import (
	"flag"
	"fmt"
)

func main() {
	word := flag.String("word", "no_translation", "Choose a word to translate")
	flag.Parse()

	translation := translate(*word)

	fmt.Println(translation)
}

func translate(word string) string {
	translations := map[string]string{
		"word":  "слово",
		"work":  "работа",
		"world": "мир",
	}

	translation := translations[word]

	if translation == "" {
		return "Translation is missing"
	} else {
		return translation
	}
}
