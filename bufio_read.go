package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	input := strings.NewReader("this is long string\nwith new line\n")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}

	reader2 := strings.NewReader("Hello, Reader!")
	buffer := make([]byte, 20)

	for {
		n, err := reader2.Read(buffer)
		if err == io.EOF {
			break
		}
		fmt.Printf("Read %d bytes: %s\n", n, buffer[:n])
	}

	reader3 := bufio.NewReader(os.Stdin)

	fmt.Print("Enter text: ")
	a, err := reader3.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println("You entered:", a)

}
