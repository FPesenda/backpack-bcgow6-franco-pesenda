package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Handler interface {
	GetAll() gin.HandlerFunc
	Get() gin.HandlerFunc
	Create() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
}

func CreateRequestTest(method string, url string, body []byte) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", os.Getenv("TOKEN"))
	return req, httptest.NewRecorder()
}

func CreateServer(handler Handler, entitypath string) *gin.Engine {
	// Print routes in console
	gin.SetMode(gin.ReleaseMode)
	//Modify token
	os.Setenv("TOKEN", "12345")
	router := gin.Default()
	router.Use(TokenAuthMiddleware())
	entityGroup := router.Group(entitypath)
	{
		entityGroup.POST("", handler.Create())
		entityGroup.GET("", handler.GetAll())
		entityGroup.GET("/:id", IdValidationMiddleware(), handler.Get())
		entityGroup.PATCH("/:id", IdValidationMiddleware(), handler.Update())
		entityGroup.DELETE("/:id", IdValidationMiddleware(), handler.Delete())
	}
	return router
}
