package main

import "fmt"

type student struct {
	name     string
	lastName string
	DNI      string
	date     string
}

func main() {
	student1 := student{
		name:     "Franco",
		lastName: "Pesenda",
		DNI:      "34543333",
		date:     "23/7/1990",
	}
	fmt.Println(student1)
}
