package main

import "github.com/gin-gonic/gin"

func main() {

	name := "Marcos"

	routerMain := gin.Default()

	routerMain.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello " + name,
		})
	})
	routerMain.Run()
}
