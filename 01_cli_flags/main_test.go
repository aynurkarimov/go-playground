package main

import (
	"os"
	"testing"
)

func TestParseFlag(t *testing.T) {
	mockedArguments := []string{"test-program", "-word=hello"}
	os.Args = mockedArguments

	output := parseFlag()

	expected := "hello"
	if output != expected {
		t.Errorf("Failed output. Expected: %q, received: %q", expected, output)
	}
}

func TestTranslateWithMissingWord(t *testing.T) {
	missingWord := "wood"

	output := translate(missingWord)

	expected := "Translation is missing"
	if output != expected {
		t.Errorf("Failed output. Expected: %q, received: %q", expected, output)
	}
}

func TestTranslateWithExistingWord(t *testing.T) {
	existingWord := "word"

	output := translate(existingWord)

	expected := "слово"
	if output != expected {
		t.Errorf("Failed output. Expected: %q, received: %q", expected, output)
	}
}
