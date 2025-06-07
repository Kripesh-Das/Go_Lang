package main

import (
	"fmt"
	"io"
	"os"
)

func main(){

	content := "he"

	file, err := os.Create("./HI")
	if err != nil{
		panic(err)
	}

	len, err := io.WriteString(file,content)

	if err!= nil{
		panic(err)
	}
	fmt.Println(len)

	// Close the file before reading it again
	file.Close()
	reader("./HI")
}	


func reader(filePath string){
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}