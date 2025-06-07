package main

import (
	"fmt"
)

func main() {
	login := 3
	var result string

	if login < 10 {
		result = "Within limit"

	}

	fmt.Println(result)
	var err error

	if err != nil {
		// handle the error here
	}
}