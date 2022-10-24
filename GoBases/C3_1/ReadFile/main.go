package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type product struct {
	Id     int     `json:"Id"`
	Price  float64 `json:"Price"`
	Amount int     `json:"Amount"`
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
	productB := product{Id: 2, Price: 21.50, Amount: 13}
	productC := product{Id: 3, Price: 5.50, Amount: 22}

	Write(productA, *fileOfProducts)
	Write(productB, *fileOfProducts)
	Write(productC, *fileOfProducts)

	fileOpen, _ := os.Open("FileOfProducts.csv")
	reader := csv.NewReader(fileOpen)

	var products []product
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		a, _ := strconv.Atoi(line[0])
		b, _ := strconv.ParseFloat(line[1], 64)
		c, _ := strconv.Atoi(line[2])
		products = append(products, product{
			Id:     a,
			Price:  b,
			Amount: c,
		})
	}
	/*
		productJSON, _ := json.Marshal(products)
		fmt.Println(string(productJSON))
	*/
	fmt.Println("ID\tPrecio\tCantidad")
	for _, v := range products {
		fmt.Println(v.Id, "\t", v.Price, "\t", v.Amount)
	}

}
