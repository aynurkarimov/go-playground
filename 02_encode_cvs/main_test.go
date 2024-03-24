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

func TestConvertLinesToProductsWithInvalidPrice(t *testing.T) {
	lines := [][]string{
		{"Banana", "2", "1000", "Grocery"},
		{"Apple", "x", "1000", "Grocery"},
	}

	_, outputError := convertLinesToProducts(lines)

	expected := "on the line (1), price must be an int"
	if outputError.Error() != expected {
		t.Errorf("Failed output. Expected: %q, received: %q", expected, outputError)
	}
}

func TestConvertLinesToProductsWithInvalidAmount(t *testing.T) {
	lines := [][]string{
		{"Banana", "2", "1000", "Grocery"},
		{"Apple", "4", "2000", "Grocery"},
		{"Apple", "6", ".", "Grocery"},
	}

	_, outputError := convertLinesToProducts(lines)

	expected := "on the line (2), amount must be an int"
	if outputError.Error() != expected {
		t.Errorf("Failed output. Expected: %q, received: %q", expected, outputError)
	}
}

func TestConvertLinesToProductsWithValidData(t *testing.T) {
	lines := [][]string{{"Banana", "2", "1000", "Grocery"}}

	output, _ := convertLinesToProducts(lines)

	expected := []Product{
		{name: "Banana", amount: 1000, price: 2, category: "Grocery"},
	}
	for i := range output {
		if reflect.DeepEqual(output[i], expected[i]) == false {
			t.Errorf("Failed output. Expected: %q, received: %q", expected, output)
		}
	}
}
