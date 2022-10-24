package main

import (
	"fmt"
	"os"
)

type product struct {
	Id     int
	Price  float64
	Amount int
}

func Write(prod product, file os.File) (err error) {
	_, errW := file.WriteString(fmt.Sprint(prod.Id, ",", prod.Price, ",", prod.Amount, "\n"))
	if errW != nil {
		err = errW
	}
	return
}

func main() {

	fileOfProducts, err := os.Create("FileOfProducts.csv")

	if err != nil {
		panic("no se pudo crear el archivo")
	}

	productA := product{Id: 1, Price: 25.50, Amount: 3}
	productB := product{Id: 2, Price: 25.50, Amount: 3}
	productC := product{Id: 3, Price: 25.50, Amount: 3}

	Write(productA, *fileOfProducts)
	Write(productB, *fileOfProducts)
	Write(productC, *fileOfProducts)
}
