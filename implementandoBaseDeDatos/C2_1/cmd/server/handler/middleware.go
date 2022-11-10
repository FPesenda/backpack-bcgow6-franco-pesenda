package handler

import (
	"net/http"
	"strconv"

	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/implementandoBaseDeDatos/C1_1/pkg/web"
	"github.com/gin-gonic/gin"
)

func IdValidationMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Abort(ctx, http.StatusBadRequest, "formate incorredto de id %s", ctx.Param("id"))
			return
		}
		ctx.Set("id_validated", id)
		ctx.Next()
	}
}
