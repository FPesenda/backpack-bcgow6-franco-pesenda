package main

import (
	"encoding/json"
	"log"
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

func main() {

	router := gin.Default()

	router.GET("/products/GetAll", func(ctx *gin.Context) {
		file, err := os.ReadFile("products.json")
		if err != nil {
			log.Fatal(err)
		}
		var products []Products
		if errJSON := json.Unmarshal(file, &products); errJSON != nil {
			log.Fatal(err)
		}
		ctx.JSON(200, gin.H{
			"DATA": products,
		})
	})

	router.Run()

}
