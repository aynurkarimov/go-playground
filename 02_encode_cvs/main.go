package main

import (
	"encoding/csv"
	"errors"
	"os"
)

func main() {
	readCsvFile("./products.csv")
}

func readCsvFile(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, errors.New("couldn't find " + filePath)
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()

	if err != nil {
		return nil, errors.New("couldn't parse " + filePath)
	}

	return lines, nil
}
