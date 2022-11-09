package handler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/pkg/web"
	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			web.Abort(c, http.StatusUnauthorized, "error does not have permissions to perform the requested request, invalid token: %s", token)
			return
		}
		c.Next()
	}
}

func IdValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Abort(c, http.StatusBadRequest, "error invalid ID: %s", c.Param("id"))
			return
		}
		c.Set("id_validated", id)
		c.Next()
	}
}
