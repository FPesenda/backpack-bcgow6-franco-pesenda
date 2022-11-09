package routes

import (
	"database/sql"
	"os"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w6-4/docs"
	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/employee"
	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/product"
	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/section"
	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/seller"
	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/warehouse"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router interface {
	MapRoutes()
}

type router struct {
	eng *gin.Engine
	rg  *gin.RouterGroup
	db  *sql.DB
}

func NewRouter(eng *gin.Engine, db *sql.DB) Router {
	return &router{eng: eng, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()

	r.buildSwaggerRoutes()
	r.buildSellerRoutes()
	r.buildProductRoutes()
	r.buildSectionRoutes()
	r.buildWarehouseRoutes()
	r.buildEmployeeRoutes()
	r.buildBuyerRoutes()
}

func (r *router) setGroup() {
	r.rg = r.eng.Group("/api/v1")
	r.rg.Use(handler.TokenAuthMiddleware())
}

func (r *router) buildSwaggerRoutes() {
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.eng.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}

func (r *router) buildSellerRoutes() {
	// sellers
	repo := seller.NewRepository(r.db)
	service := seller.NewService(repo)
	sellerHandler := handler.NewSeller(service)
	r.rg.GET("/sellers", sellerHandler.GetAll())
	r.rg.POST("/sellers", sellerHandler.Create())
	r.rg.GET("/sellers/:id", handler.IdValidationMiddleware(), sellerHandler.Get())
	r.rg.PATCH("/sellers/:id", handler.IdValidationMiddleware(), sellerHandler.Update())
	r.rg.DELETE("/sellers/:id", handler.IdValidationMiddleware(), sellerHandler.Delete())
}

func (r *router) buildProductRoutes() {
	repository := product.NewRepository(r.db)
	service := product.NewService(repository)
	productHandler := handler.NewProduct(service)
	productGroup := r.rg.Group("/products")
	{
		productGroup.POST("/", productHandler.Create())
		productGroup.GET("/", productHandler.GetAll())
		productGroup.GET("/:id", handler.IdValidationMiddleware(), productHandler.Get())
		productGroup.PATCH("/:id", handler.IdValidationMiddleware(), productHandler.Update())
		productGroup.DELETE("/:id", handler.IdValidationMiddleware(), productHandler.Delete())
	}
}

func (r *router) buildSectionRoutes() {
	repo := section.NewRepository(r.db)
	service := section.NewService(repo)
	section := handler.NewSection(service)
	sectionGroup := r.rg.Group("/sections")
	{
		sectionGroup.POST("/", section.Create())
		sectionGroup.GET("/", section.GetAll())
		sectionGroup.GET("/:id", handler.IdValidationMiddleware(), section.Get())
		sectionGroup.PATCH("/:id", handler.IdValidationMiddleware(), section.Update())
		sectionGroup.DELETE("/:id", handler.IdValidationMiddleware(), section.Delete())
	}
}

func (r *router) buildWarehouseRoutes() {
	repo := warehouse.NewRepository(r.db)
	service := warehouse.NewService(repo)
	warehouseHandler := handler.NewWarehouse(service)
	whGroup := r.rg.Group("/warehouses")
	{
		whGroup.POST("/", warehouseHandler.Create())
		whGroup.GET("/", warehouseHandler.GetAll())
		whGroup.GET("/:id", handler.IdValidationMiddleware(), warehouseHandler.Get())
		whGroup.PATCH("/:id", handler.IdValidationMiddleware(), warehouseHandler.Update())
		whGroup.DELETE("/:id", handler.IdValidationMiddleware(), warehouseHandler.Delete())
	}

}

func (r *router) buildEmployeeRoutes() {
	// Example
	repo := employee.NewRepository(r.db)
	service := employee.NewService(repo)
	handler := handler.NewEmployee(service)
	r.rg.POST("/employees", handler.Create())
	r.rg.GET("/employees", handler.GetAll())
	r.rg.GET("/employees/:id", handler.Get())
	r.rg.PATCH("employees/:id", handler.Update())
	r.rg.DELETE("/employees/:id", handler.Delete())
}

func (r *router) buildBuyerRoutes() {
	// Example
	repo := buyer.NewRepository(r.db)
	service := buyer.NewService(repo)
	buyer := handler.NewBuyer(service)
	r.rg.POST("/buyers", buyer.Create())
	r.rg.GET("/buyers", buyer.GetAll())
	r.rg.GET("/buyers/:id", handler.IdValidationMiddleware(), buyer.Get())
	r.rg.DELETE("/buyers/:id", handler.IdValidationMiddleware(), buyer.Delete())
	r.rg.PATCH("/buyers/:id", handler.IdValidationMiddleware(), buyer.Update())

}
