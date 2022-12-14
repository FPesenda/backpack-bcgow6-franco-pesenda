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
	GetAll() ([]domains.Product, error)
	Delete(id int) error
	Exists(id int) bool
	Update(domains.Product) error
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

type repository struct {
	db *sql.DB
}

var (
	SELECT_BY_NAME = "SELECT id , name , type , count , price FROM products WHERE name =?;"
	SELECT_ALL     = "SELECT id , name , type , count , price FROM products ;"
	INSERT_PRODUCT = "INSERT INTO products (name, type, count, price) VALUES (?,?,?,?);"
	DELETE_BY_ID   = "DELETE FROM products WHERE id =?;"
	UPDATE         = "UPDATE products SET name=?, type=?, count=?, price=? WHERE id=?;"
	EXIST          = "SELECT id FROM products WHERE id=?"
)

var (
	db_error_unprocesableEntity = errors.New("No se pudo procesar la entidad")
	db_error_failtConsult       = errors.New("Error creando la consulta")
	db_error_notFound           = errors.New("Prodcto no encontrado")
	db_error_fail               = errors.New("Db error fatal")
	consult_error_no_afected    = errors.New("Consult no afect")
)

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

/*
Ejercicio 2 - Implementar GetAll()
Dise??ar un m??todo GetAll.
Dentro del archivo repository desarrollar el m??todo GetAll().
Comprobar el correcto funcionamiento.
*/

func (r *repository) GetAll() (products []domains.Product, err error) {

	statement, err := r.db.Prepare(SELECT_ALL)

	if err != nil {
		err = db_error_failtConsult
		return
	}

	defer statement.Close()

	rows, err := statement.Query()

	for rows.Next() {

		var productTemporal domains.Product

		err = rows.Scan(
			&productTemporal.ID,
			&productTemporal.Name,
			&productTemporal.Type,
			&productTemporal.Count,
			&productTemporal.Price,
		)

		if err != nil {
			err = db_error_unprocesableEntity
			products = nil
			return
		}

		products = append(products, productTemporal)

	}

	err = rows.Err()

	if err != nil {
		err = errors.New("Error en la iteraci??n")
		products = nil
		return
	}

	return
}

/*
Ejercicio 3 - Implementar Delete()
Dise??ar un m??todo para eliminar un recurso de la base de datos.
Dentro del archivo repository desarrollar el m??todo Delete().
Comprobar el correcto funcionamiento.
*/

func (r *repository) Delete(id int) (err error) {

	statement, err := r.db.Prepare(DELETE_BY_ID)

	if err != nil {
		err = db_error_failtConsult
		return
	}

	defer statement.Close()

	_, err = statement.Exec(id)

	if err != nil {
		err = db_error_notFound
		return
	}
	return
}

func (r *repository) Update(product domains.Product) (err error) {
	statement, err := r.db.Prepare(UPDATE)

	if err != nil {
		err = db_error_failtConsult
		return
	}

	result, err := statement.Exec(
		product.Name,
		product.Type,
		product.Count,
		product.Price,
		product.ID,
	)

	if err != nil {
		err = db_error_unprocesableEntity
		return
	}

	rowsAfected, err := result.RowsAffected()

	if err != nil {
		return
	}

	if rowsAfected < 1 {
		err = consult_error_no_afected
	}
	return
}

func (r *repository) Exists(id int) bool {
	statement, err := r.db.Prepare(EXIST)

	if err != nil {
		panic(err)
	}

	err = statement.QueryRow(id).Scan(&id)

	return err == nil
}
