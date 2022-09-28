package main

import (
	"fmt"
	"go/printer"
	"os"
)

type Products struct {
	Id   string
	Price  float64
	Amount int
}

func writte (product Products) (line string){
	line := fmt.Sprint(product.Id,",",product.Price,",",product.Amount)
}

func main() {
	fileOfProducts, err := os.Create("FileOfProducts.txt")
	
	productA := Products{Id: 1, Price: 25.50, Amount: 3}
	productB := Products{Id: 2, Price: 25.50, Amount: 3}
	productC := Products{Id: 3, Price: 25.50, Amount: 3}

	toWritte string

}
