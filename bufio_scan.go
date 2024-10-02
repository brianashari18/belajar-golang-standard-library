package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter text (Ctrl+D to end):")
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("You entered:", text)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

	// scanner.Scan()
	// input := scanner.Text()
	// fmt.Println(input)
}
