package main

import "fmt"

func main(){
	fmt.Println("Hi")

	user := User{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}
	fmt.Println("User email:", user.Email)
	fmt.Println("User age:", user.Age)
	fmt.Println("User name:", user.Name)
	fmt.Println(user)

	user.status()
}

type User struct {
	Name string
	Age int 
	Email string
}

func (u User) status(){
	fmt.Println(u.Email)
}