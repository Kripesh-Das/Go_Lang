package main

import "fmt"

func main(){

	days := []string{"Mon","Tue","Thu","Fri","Sat"}
	fmt.Println(days)
	
	for d:=0; d<len(days);d++{
		fmt.Println(days[d])
	}
    // ......................................
	for d:=range days{
		fmt.Println(days[d])
	}
	// ......................................
	for index,day := range days{
		fmt.Printf("Index is %v for the Day %v \n",index,day)
	}
	// ......................................
	var i int = 1
	for i<10{
		i++
	}
}
