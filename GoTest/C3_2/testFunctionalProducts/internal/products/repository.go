package products

import (
	"errors"
	"fmt"

	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/GOWeb/C3_2/AddStore/pkg/store"
	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/GoTest/C3_2/testFunctionalProducts/internal/domain"
)

type Products struct {
	Id    int
	Name  string  `binding:"required"`
	Color string  `binding:"required"`
	Code  string  `binding:"required"`
	Price float64 `binding:"required"`
}

type repository struct {
	db store.Store
}

var lastId int

type Repository interface {
	Patch(int, string, float64) (Products, error)
	Delete(int) (int, error)
	GetAll() ([]Products, error)
	Store(int, string, string, string, float64) (Products, error)
	UpdateByID(int, string, string, string, float64) (Products, error)
	GetLastId() (int, error)
	ChangeName(int, string) (Products, error)
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func readDB(rep *repository) (data []Products, err error) {
	err = rep.db.Read(&data)
	return
}

func writeDB(rep *repository, data []Products) (err error) {
	err = rep.db.Write(data)
	return
}

func (rep *repository) Patch(id int, name string, price float64) (product Products, err error) {
	var productTemporal Products
	update := false
	ProductsBBDD, err := readDB(rep)
	for i := range ProductsBBDD {
		if ProductsBBDD[i].Id == id {
			ProductsBBDD[i].Name = name
			ProductsBBDD[i].Price = price
			update = true
			productTemporal = ProductsBBDD[i]
		}
	}
	if !update {
		err = errors.New(fmt.Sprint("No se encontro el elemento ", id, " en la BBDD"))
		return
	}
	product = productTemporal
	errw := writeDB(rep, ProductsBBDD)
	if errw != nil {
		err = errw
		return
	}
	return
}

func (rep *repository) ChangeName(id int, name string) (product Products, err error) {
	update := false
	ProductsBBDD, err := readDB(rep)
	for i := range ProductsBBDD {
		if ProductsBBDD[i].Id == id {
			update = true
			ProductsBBDD[i].Name = name
			product = ProductsBBDD[i]
			errw := writeDB(rep, ProductsBBDD)
			if errw != nil {
				err = errw
				return
			}
			break
		}
	}
	if !update {
		err = errors.New(fmt.Sprint("No se encontro el elemento ", id, " en la BBDD"))
		return
	}
	return
}

func (rep *repository) Delete(id int) (deleted int, err error) {
	deleted = 0
	var index int
	ProductsBBDD, errR := readDB(rep)
	if errR != nil {
		err = errR
	}
	for i := range ProductsBBDD {
		if ProductsBBDD[i].Id == id {
			deleted = 1
			index = i
			break
		}
	}
	if deleted == 0 {
		err = fmt.Errorf("No se encuentra el elemento que se quiere borrar con id %d", id)
		return
	}
	ProductsBBDD = append(ProductsBBDD[:index], ProductsBBDD[index+1:]...)
	errw := writeDB(rep, ProductsBBDD)
	if errw != nil {
		err = errw
		return
	}
	return
}

func (rep *repository) GetAll() ([]Products, error) {
	ProductsBBDD, errR := readDB(rep)
	if errR != nil {
		return nil, errR
	}
	if len(ProductsBBDD) == 0 {
		return nil, fmt.Errorf("No hay elementos en la base de datos")
	}
	return ProductsBBDD, nil
}

func (rep *repository) Store(id int, name, color, code string, price float64) (p Products, err error) {
	ProductsBBDD, errR := readDB(rep)
	if errR != nil {
		err = errR
		return
	}
	p = Products{id, name, color, code, price}
	ProductsBBDD = append(ProductsBBDD, p)
	lastId = p.Id
	errw := writeDB(rep, ProductsBBDD)
	if errw != nil {
		err = errw
		return
	}
	return p, nil
}

func (rep *repository) UpdateByID(id int, name, color, code string, price float64) (prod domain.Product, err error) {
	productTemporal := Products{
		Id:    id,
		Name:  name,
		Color: color,
		Code:  code,
		Price: price,
	}
	updated := false
	ProductsBBDD, errR := readDB(rep)
	if errR != nil {
		err = errR
	}
	for i := range ProductsBBDD {
		if ProductsBBDD[i].Id == id {
			ProductsBBDD[i] = productTemporal
			updated = true
		}
	}
	if !updated {
		err = errors.New(fmt.Sprint("No se pudo encontrar el producto con id ", id))
		return
	}
	prod = productTemporal
	errw := writeDB(rep, ProductsBBDD)
	if errw != nil {
		err = errw
		return
	}
	return
}

func (rep *repository) GetLastId() (id int, err error) {
	ProductsBBDD, errR := readDB(rep)
	if errR != nil {
		err = errR
	}
	if len(ProductsBBDD) == 0 {
		id = 0
		return
	}
	id = lastId
	return
}
