package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type products struct {
	Id          int
	Stock       int
	Name        string    `binding:"required"`
	Color       string    `binding:"required"`
	Code        string    `binding:"required"`
	Price       float64   `binding:"required"`
	Publish     bool      `binding:"required"`
	DateCreatin time.Time `binding:"required"`
}

func addProduct(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	var productsDB []products

	file, err := os.ReadFile("products.json")

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "no se pudo acceder a la Base de Datos"})
		return
	}

	if errRead := json.Unmarshal(file, productsDB); errRead != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "no se pudo acceder a la Base de Datos"})
		return
	}

	if token != "1234" {
		ctx.JSON(http.StatusNonAuthoritativeInfo, gin.H{"error": "Token invalido"})
		return
	}

	var productIn products

	if err := ctx.ShouldBindJSON(&productIn); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productIn.Id = (len(productsDB) + 1)
	productsDB = append(productsDB, productIn)
	saveBD(productIn)
	ctx.JSON(http.StatusOK, gin.H{"status": "se guardo el producto nuevo"})
}

// SAVE IN BBBDD
func saveBD(pnew products) (err error) {

	var productsDB []products

	fileOpen, err := os.OpenFile("products.json", os.O_RDWR, 0644)
	fileRead, err := os.ReadFile("products.json")
	if err != nil {
		err = errors.New("no se pudo leer el archivo")
		return
	}

	if errRead := json.Unmarshal(fileRead, &productsDB); errRead != nil {
		err = errors.New("No se pudo transfomar el archivo a Productos")
		return
	}

	productsDB = append(productsDB, pnew)
	data, _ := json.Marshal(productsDB)
	if _, errorWrite := fileOpen.WriteString(string(data)); errorWrite != nil {
		err = errors.New("No se pudo guardar en la BBDD")
	}

	fmt.Print(productsDB)
	return
}

// SHOW BBD
func showBBDD() (err error) {
	var productsBD []products
	file, errRead := os.ReadFile("productos.json")
	if errRead != nil {
		err = errors.New("No se pudo leer el archivo")
	}
	if errjson := json.Unmarshal(file, &productsBD); errjson != nil {
		err = errors.New("no se pudo convertir la BBDD a la memoria volatil")
	}
	for _, v := range productsBD {
		fmt.Println(v)
	}
	return
}

func main() {

	router := gin.Default()

	productsGroup := router.Group("/products")
	{
		productsGroup.POST("/", addProduct)
	}
	/*
		date, errFormat := time.Parse("2006-01-02", "2018-01-20")
		if errFormat != nil {
			log.Fatal("No se pudo parsear")
		}
		p := products{
			Id:          1,
			Stock:       10,
			Name:        "Mouse",
			Color:       "Black",
			Code:        "Ms1",
			Price:       10.40,
			Publish:     true,
			DateCreatin: date,
		}
	*/
	router.Run()
}
