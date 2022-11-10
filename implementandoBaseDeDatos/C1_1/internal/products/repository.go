package products

//github.com/bootcamp-go/storage
import (
	"database/sql"
	"errors"

	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/implementandoBaseDeDatos/C1_1/internal/domains"
)

type Repository interface {
	Store(p domains.Product) (int, error)
	GetByName(name string) (domains.Product, error)
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

type repository struct {
	db *sql.DB
}

var (
	SELECT_BY_NAME = "SELECT id , name , type , count , price FROM products WHERE name =?;"
	INSERT_PRODUCT = "INSERT INTO products (name, type, count, price) VALUES (?,?,?,?);"
)

var (
	db_error_unprocesableEntity = errors.New("No se pudo procesar la entidad")
	db_error_failtConsult       = errors.New("Error creando la consulta")
	db_error_notFound           = errors.New("Prodcto no encontrado")
)

/*
Ejercicio 1 - Implementar GetByName()
Desarrollar un método en el repositorio que permita hacer búsquedas de un producto por nombre.
Para lograrlo se deberá:Diseñar interfaz “Repository” en la que exista un método GetByName()
que reciba por parámetro un string y retorne un objeto del tipo Product.
Implementar el método de forma que con el string recibido lo use para buscar en la DB por el
campo “name”.
*/

func (r *repository) GetByName(name string) (product domains.Product, err error) {

	statement, err := r.db.Prepare(SELECT_BY_NAME) // => creo la query antes de ejecutarla

	if err != nil {
		err = db_error_failtConsult
		return
	}

	defer statement.Close()

	err = statement.QueryRow(name).Scan(
		&product.ID,
		&product.Name,
		&product.Type,
		&product.Count,
		&product.Price)

	if err != nil {
		err = db_error_notFound
		return
	}
	return
}

/*
Ejercicio 2 - Replicar Store()
Tomar el ejemplo visto en la clase y diseñar el método Store():
Puede tomar de ejemplo la definición del método Store visto en clase para incorporarlo en la interfaz.
Implementar el método Store.
*/

func (r *repository) Store(product domains.Product) (lastId int, err error) {

	statement, err := r.db.Prepare(INSERT_PRODUCT)

	if err != nil {
		err = db_error_failtConsult
		lastId = -1
		return
	}

	defer statement.Close()

	result, err := statement.Exec(
		product.Name,
		product.Type,
		product.Count,
		product.Price)

	if err != nil {
		err = db_error_unprocesableEntity
		lastId = -1
		return
	}

	lastId64, err := result.LastInsertId()
	if err != nil {
		err = errors.New("Error al obtener el Id")
		lastId = -1
		return
	}

	lastId = int(lastId64)
	return
}
