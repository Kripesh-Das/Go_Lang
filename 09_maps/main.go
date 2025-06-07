package main

import "fmt"

func main(){

	lang := make(map[string]string)
	
	lang["JS"] = "Javascript"
	lang["TS"] = "Typescript"
	lang["RB"] = "Ruby"

	fmt.Println(lang)
	fmt.Println(lang["RB"])

	delete(lang,"RB")
	fmt.Println(lang)

	//............................

	for key,value := range lang {
		fmt.Printf("value is %v for the key %v \n", value, key)
	}
}	