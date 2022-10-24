package main

import "fmt"

type market struct {
	listProducts []Products
}

type products struct {
	TypeProduct string
	Name        string
	Price       float64
}

func (prod *products) CostCalculate() (cost float64) {
	cost = prod.Price
	switch prod.TypeProduct {
	case "mediano":
		cost = (prod.Price * 1.03)
	case "grande":
		cost = ((prod.Price * 1.06) + 2500.0)
	}
	return
}

type Products interface {
	CostCalculate() float64
}

type Ecommerce interface {
	total() float64
	add(Products)
}

func (mark market) total() (t float64) {
	t = 0.0
	for _, v := range mark.listProducts {
		t += v.CostCalculate()
	}
	return
}

func (mark *market) add(prod Products) {
	mark.listProducts = append(mark.listProducts, prod)
}

func (mark *market) NewEcommerce() Ecommerce {
	return &market{}
}

func (prod *products) NewProducts(tp, namep string, prc float64) Products {
	return &products{TypeProduct: tp, Name: namep, Price: prc}
}

func main() {

	ecommerce1 := market{}
	product1 := products{
		TypeProduct: "grande",
		Name:        "hojas",
		Price:       11.23,
	}
	product2 := products{
		TypeProduct: "mediano",
		Name:        "mouse",
		Price:       31.23,
	}
	product3 := products{
		TypeProduct: "chico",
		Name:        "cooler",
		Price:       21.23,
	}
	product4 := products{
		TypeProduct: "mediano",
		Name:        "McBoock",
		Price:       950.10,
	}
	ecommerce1.add(&product1)
	ecommerce1.add(&product2)
	ecommerce1.add(&product3)
	ecommerce1.add(&product4)
	fmt.Println(len(ecommerce1.listProducts))
	fmt.Println(product1.CostCalculate())
	fmt.Println(product2.CostCalculate())
	fmt.Println(product3.CostCalculate())
	fmt.Println(product4.CostCalculate())

	fmt.Println(ecommerce1.total())

}
