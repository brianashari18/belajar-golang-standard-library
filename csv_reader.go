package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "brian,anashari,\n" + "indah,cahya,nabilla\n" +
		"cornelia,vanisa,\n" + "gita,sekar,andarini"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}
