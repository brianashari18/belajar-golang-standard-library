package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Brian Anashari", "Brian"))
	fmt.Println(strings.Split("Brian Anashari", " "))
	fmt.Println(strings.ToLower("Brian Anashari"))
	fmt.Println(strings.ToUpper("Brian Anashari"))
	fmt.Println(strings.Trim("           Brian Anashari            ", " "))
	fmt.Println(strings.ReplaceAll("Satu tambah Satu", "Satu", "Dua"))
}
