package main

import (
	"fmt"
)

func main() {
	var decimal int
	var temp string

	fmt.Scan(&decimal)
	for decimal > 0 {
		temp += string(decimal % 2)
		decimal /= 2
	}

	for i := len(temp) - 1; i >= 0; i-- {
		fmt.Printf("%s", string(temp[i]))
	}
}
