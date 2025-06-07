package main

import "fmt"

func main(){
	var fruits[3]string

	fruits [0] = "apple"
	fruits [2] = "banana"

	fmt.Println("fruits:",fruits)
	fmt.Println("length", len(fruits))
}