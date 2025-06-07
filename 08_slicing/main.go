package main

import (
	"fmt"
	"sort"
)

func main(){

	var fruits = []string {"apple","banana","mango"}
	fruits = append(fruits,"guava")
	fmt.Println(fruits)
// .....................................
	fruits = append(fruits[1:],"dragon fruit")
	fmt.Println(fruits)
// .....................................

	score := make([]int,4)
	score [0] = 1
	score [1] = 11
	score [2] = 111
	score [3] = 1111

	score = append(score,11111)
	fmt.Println(score)
// .....................................

	sort.Ints(score)
	fmt.Println(score)
// .....................................

var obj = []string{"bat","ball","stumps","bails","helmet","gloves"} 
var index int = 2
fmt.Println(obj)
obj = append(obj[:index],obj[index+1:]... )
fmt.Println(obj)

}