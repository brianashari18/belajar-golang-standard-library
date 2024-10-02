package main

import (
	"fmt"
	"regexp"
)

type Item struct {
	Name string
}

func main() {
	items := []Item{
		{"Xiaomi 13T"},
		{"Xiaomi 12T"},
		{"Iphone 15 Pro Max"},
		{"Iphone 15 Pro"},
		{"Iphone 15"},
		{"Samsung Galaxy S24 Ultra"},
		{"Samsung Galaxy S24 Plus"},
		{"Samsung Galaxy S24"},
		{"Samsung Galaxy S23"},
		{"Samsung Galaxy S22"},
	}

	var regex *regexp.Regexp = regexp.MustCompile(`a([a-z])a`)

	fmt.Println(regex.MatchString("aja"))
	fmt.Println(regex.MatchString("apa"))
	fmt.Println(regex.MatchString("aDa"))

	fmt.Println(regex.FindAllString("aja apa ada aka asa alat Apa a1a", 2))

	input := "Samsung Galaxy S23"
	matchedItems := []string{}
	var regexInput = regexp.MustCompile(regexp.QuoteMeta(input))
	for _, value := range items {
		if regexInput.MatchString(value.Name) {
			matchedItems = append(matchedItems, value.Name)
		}
	}

	fmt.Println(matchedItems)
}
