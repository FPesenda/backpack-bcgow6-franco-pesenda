package main

import "fmt"

type user struct {
	Name     string
	LastName string
	Email    string
	Products []product
}

type product struct {
	Name     string
	Price    float64
	Quantity int
}

func (prod *product) newProduct(name string, price float64) (productNew product) {
	productNew = product{
		Name:  name,
		Price: price,
	}
	return
}

func (usr *user) addProuct(prod product) {
	usr.Products = append(usr.Products, prod)
}

func (usr *user) deleteProducts() {
	var emptySlice []product
	usr.Products = emptySlice
}

func main() {
	product1 := product{
		Name:     "producto1",
		Price:    34.34,
		Quantity: 10,
	}
	product2 := product{
		Name:     "producto2",
		Price:    31.31,
		Quantity: 5,
	}
	user1 := user{
		Name:     "Franco",
		LastName: "Pesenda",
		Email:    "franco@gmail.com",
	}
	user1.addProuct(product1)
	user1.addProuct(product2)
	fmt.Println(user1.Products)
	user1.deleteProducts()
	fmt.Println(user1.Products)
}
