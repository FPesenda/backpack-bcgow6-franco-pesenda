/*
Los productos varían por id, nombre, color, precio, stock, código (alfanumérico), publicado (si-no), fecha de creación.
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Products struct {
	Id          int
	Stock       int
	Name        string
	Color       string
	Code        string
	Price       float64
	Publish     bool
	DateCreatin time.Time
}

func main() {

	var MyProducts []Products

	fl, err := os.ReadFile("products.json")
	if err != nil {
		panic("No se pudo leer el archivo")
	}
	if errJSON := json.Unmarshal(fl, &MyProducts); errJSON != nil {
		log.Fatal(errJSON)
	}
	fmt.Print("Productos: ", MyProducts)
}
