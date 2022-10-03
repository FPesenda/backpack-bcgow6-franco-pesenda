package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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

func getById(cxt *gin.Context) {

	ID, _ := strconv.Atoi(cxt.Param("id"))
	var products []Products

	file, err := os.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}
	if errJSON := json.Unmarshal(file, &products); errJSON != nil {
		log.Fatal(err)
	}

	for _, v := range products {
		if v.Id == ID {
			cxt.JSON(http.StatusOK, v)
		}
	}
	cxt.JSON(http.StatusNotFound, gin.H{"status": fmt.Sprint("El producto con Id ", ID, " no se encuentra en la base de datos")})
}

func main() {

	router := gin.Default()

	products := router.Group("/products")
	{
		products.GET("/getById/:id", getById)
	}

	router.Run()

}
