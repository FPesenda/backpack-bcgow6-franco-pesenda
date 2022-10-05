package products

import (
	"fmt"
)

type Products struct {
	Id    int
	Name  string  `binding:"required"`
	Color string  `binding:"required"`
	Code  string  `binding:"required"`
	Price float64 `binding:"required"`
}

type repository struct {
}

var ProductsBBDD []Products
var lastId int

type Repository interface {
	GetAll() ([]Products, error)
	Store(int, string, string, string, float64) (Products, error)
	UpdateByID(int, string, string, string, float64) (Products, error)
	GetLastId() (int, error)
}

func NewRepository() Repository {
	return &repository{}
}

func (rep *repository) GetAll() ([]Products, error) {
	if len(ProductsBBDD) == 0 {
		return nil, fmt.Errorf("No hay elementos en la base de datos")
	}
	return ProductsBBDD, nil
}

func (rep *repository) Store(id int, name, color, code string, price float64) (Products, error) {
	p := Products{id, name, color, code, price}
	ProductsBBDD = append(ProductsBBDD, p)
	lastId = p.Id
	return p, nil
}

func (rep *repository) UpdateByID(id int, name, color, code string, price float64) (prod Products, err error) {
	productTemporal := Products{
		Id:    id,
		Name:  name,
		Color: color,
		Code:  code,
		Price: price,
	}
	updated := false
	for i, _ := range ProductsBBDD {
		if ProductsBBDD[i].Id == id {
			ProductsBBDD[i] = productTemporal
			updated = true
		}
	}
	if !updated {
		err = fmt.Errorf("No se pudo encontrar el producto con id %d", id)
		return
	}
	prod = productTemporal
	return
}

func (rep *repository) GetLastId() (id int, err error) {
	if len(ProductsBBDD) == 0 {
		id = 0
		return
	}
	id = lastId
	return
}
