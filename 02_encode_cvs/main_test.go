package main

import (
	"encoding/csv"
	"os"
	"reflect"
	"testing"
)

func TestReadCsvFileWithValidFile(t *testing.T) {
	file, err := os.Create("_valid.csv")

	if err != nil {
		t.Errorf("Couldn't initialize test")
	}

	defer file.Close()
	defer os.Remove("_valid.csv")

	mockedCsvLines := [][]string{{"Product", "Price"}, {"Apple", "10"}, {"Tomato", "42"}}
	writer := csv.NewWriter(file)
	writer.WriteAll(mockedCsvLines)
	writer.Flush()

	output, _ := readCsvFile("_valid.csv")

	expected := [][]string{{"Product", "Price"}, {"Apple", "10"}, {"Tomato", "42"}}
	if reflect.DeepEqual(output, expected) == false {
		t.Errorf("Failed output. Expected: %q, received: %q", expected, output)
	}
}

func TestReadCsvFileWithInvalidPath(t *testing.T) {
	_, outputError := readCsvFile("_invalid_path.csv")

	expected := "couldn't find _invalid_path.csv"
	if outputError.Error() != expected {
		t.Errorf("Failed output. Expected: %q, received: %q", expected, outputError)
	}
}
