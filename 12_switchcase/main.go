package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){

	var dice int 
	r := rand.New(rand.NewSource(time.Now().UnixNano())) 
	dice = r.Intn(6) + 1

	fmt.Println("Dice rolled:", dice)

	switch dice{
	case 1:
		fmt.Println("hi")
	case 2:
		fmt.Println("hi")
	case 3:
		fmt.Println("hi")
	case 4:
		fmt.Println("hi")
	case 5:
		fmt.Println("hi")
		fallthrough
	case 6:
		fmt.Println("hi")

	default:
		fmt.Print("")
	}
}