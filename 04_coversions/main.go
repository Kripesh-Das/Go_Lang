package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter number")
	input,_ := reader.ReadString('\n')
	
	rating, error := strconv.ParseFloat(strings.TrimSpace(input),64)

	if error != nil {
		
		fmt.Println(error)
	}else{
	
		fmt.Println("Adding +1 to your rating:", rating+1)
	}
}