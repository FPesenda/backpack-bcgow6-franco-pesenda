package main

import "fmt"

type User struct {
	Name     string
	LastName string
	Age      int
	Email    string
	Password []string
}

var pointerString = new(string)
var pointerString2 *string

func main() {

	myUser := User{
		Name:     "Franco",
		LastName: "Pesenda",
		Email:    "miEmail@ml.com",
	}

	name := "Marcelo"
	lastName := "Parilli"
	pointerString = &name
	pointerString2 = &lastName
	myUser.changeNameAndLastName(*pointerString, *pointerString2)

	fmt.Println(myUser)
}
