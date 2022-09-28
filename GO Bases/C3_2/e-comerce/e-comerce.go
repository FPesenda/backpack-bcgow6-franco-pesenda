package main

import "fmt"

type User struct {
	Name     string
	LastName string
	Email    string
	Products []Products
}

type Products struct {
	Name     string
	Price    float64
	Quantity int
}

func (product *Products) newProduct(name string, price float64) (productNew Products) {
	productNew = Products{
		Name:  name,
		Price: price,
	}
}

func main() {

}
