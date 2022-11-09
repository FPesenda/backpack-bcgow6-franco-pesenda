package main

import (
	"database/sql"

	"log"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/cmd/server/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// @BasePath /api/v1

// @title          Grupo 4 - Proyecto Integrador
// @version        1.0
// @description    This API Handle MELI - Fresh products
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url  https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}

	// NO MODIFICAR
	db, err := sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123@/melisprint")
	if err != nil {
		panic(err)
	}

	eng := gin.Default()

	router := routes.NewRouter(eng, db)
	router.MapRoutes()

	if err := eng.Run(); err != nil {
		panic(err)
	}

}
