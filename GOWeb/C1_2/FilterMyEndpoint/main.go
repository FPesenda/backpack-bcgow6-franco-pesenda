package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
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

func getByColor(cxt *gin.Context) {

	ColorFilter := cxt.Param("color")
	var products []Products

	file, err := os.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}
	if errJSON := json.Unmarshal(file, &products); errJSON != nil {
		log.Fatal(err)
	}

	for _, v := range products {
		if v.Color == ColorFilter {
			cxt.JSON(http.StatusOK, v)
			return
		}
	}
	cxt.JSON(http.StatusNotFound, gin.H{"status": fmt.Sprint("Producto con la categoria ", ColorFilter, " no se encuentra en la base de datos")})
}

func getAll(cxt *gin.Context) {
	file, err := os.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}
	var products []Products
	if errJSON := json.Unmarshal(file, &products); errJSON != nil {
		log.Fatal(err)
	}
	cxt.JSON(http.StatusOK, products)
}

func main() {

	router := gin.Default()

	products := router.Group("/products")
	{
		products.GET("/getByColor/:color", getByColor)
		products.GET("/", getAll)
	}

	router.Run()

}
