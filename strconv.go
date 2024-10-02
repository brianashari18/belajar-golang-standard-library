package main

import (
	"fmt"
	"strconv"
)

func main() {
	if resultBool, err := strconv.ParseBool("s"); err == nil {
		fmt.Println(resultBool)
	} else {
		fmt.Println("Error: ", err.Error())
	}

	if resultInt, err := strconv.ParseInt("10", 10, 10); err == nil {
		fmt.Println(resultInt)
	} else {
		fmt.Println("Error: ", err.Error())
	}

	formatInt := strconv.Itoa(1203912391)
	fmt.Println(formatInt)

}
