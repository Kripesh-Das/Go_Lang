package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter anything")
	input,_ := reader.ReadString('\n')
	fmt.Printf("%T",input)
	
}