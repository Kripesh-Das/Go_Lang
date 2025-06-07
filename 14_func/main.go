package main

import "fmt"

func main(){
	result := adder(3, 5)
	fmt.Println(result)

	r := pro(2,3,4,5)
	fmt.Println(r)
}

func adder(val1 int, val2 int) int{
	return val1 + val2
}

func pro(values ...int) int{
	var tot int
	for _,val := range values{
		tot = tot + val
	}

	return tot
}