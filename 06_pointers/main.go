package main

import "fmt"

func main(){

	var ptr *int 
	fmt.Println(ptr)

	var num int = 23
	var point = &num

	fmt.Println(point)
	fmt.Println(*point)

	*point = *point + 2
	fmt.Println(*point)
}