package products

import "errors"

type products struct {
	Id    int
	Stock int
	Name  string  `binding:"required"`
	Color string  `binding:"required"`
	Code  string  `binding:"required"`
	Price float64 `binding:"required"`
}

type repository struct {
}

var productsBBDD []products
var lastId int

type Repository interface {
	GetById(int) (products, error)
	GetAll() ([]products, error)
	Save(products) (products, error)
	DeleteById(int) (int, error)
	GetLastId() (int, error)
}

func NewRepository() Repository {
	return &repository{}
}

func (rep *repository) GetById(id int) (product products, err error) {
	for _, v := range productsBBDD {
		if v.Id == id {
			product = v
			return
		}
	}
	err = errors.New("No se encintro el revurso con el Id especificado")
	return
}

func (rep *repository) GetAll() (products []products, err error) {
	if len(products) == 0 {
		err = errors.New("BBDD vacia")
		return
	}
	for _, v := range productsBBDD {
		products = append(products, v)
	}
	return
}

func (rep *repository) Save(product products) (prod products, err error) {
	productsBBDD = append(productsBBDD, product)
	lastId = product.Id
	prod = product
	return
}

func (rep *repository) DeleteById(id int) (deleted int, err error) {
	var index int
	for i, v := range productsBBDD {
		if v.Id == id {
			index = i
		}
	}
	productsBBDD[index] = productsBBDD[len(productsBBDD)]
	//productsBBDD[len(productsBBDD)] = nil
	productsBBDD = productsBBDD[:len(productsBBDD)-1]
	deleted = 1
	return
}

func (rep *repository) GetLastId() (id int, err error) {
	if len(productsBBDD) == 0 {
		err = errors.New("No hay Id en la BBDD")
	}
	id = lastId
	return
}
