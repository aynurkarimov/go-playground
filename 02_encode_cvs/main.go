package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	lines, err := readCsvFile("./products.csv")

	if err != nil {
		fmt.Println(err)

		os.Exit(0)
	}

	products, err := convertLinesToProducts(lines)

	if err != nil {
		fmt.Println(err)

		os.Exit(0)
	}

	for _, product := range products {
		fmt.Println(product.name, product.PricierThanTen())
	}
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

type Product struct {
	name     string
	price    int
	amount   int
	category string
}

func (product Product) PricierThanTen() bool {
	return product.price > 10
}

func convertLinesToProducts(lines [][]string) ([]Product, error) {
	products := make([]Product, len(lines))

	for idx, line := range lines {
		convertedPrice, err := strconv.Atoi(line[1])

		if err != nil {
			return nil, errors.New("on the line (" + strconv.Itoa(idx) + "), price must be an int")
		}

		convertedAmount, err := strconv.Atoi(line[2])

		if err != nil {
			return nil, errors.New("on the line (" + strconv.Itoa(idx) + "), amount must be an int")
		}

		products[idx] = Product{
			name:     line[0],
			price:    convertedPrice,
			amount:   convertedAmount,
			category: line[3],
		}
	}

	return products, nil
}
