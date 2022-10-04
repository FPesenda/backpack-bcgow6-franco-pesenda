package main

type service struct {
	repository Repository
}

type Service interface {
}

func NewService() Service {
	return &service{}
}

func main() {

}

func (serv *service) GetAll() (products []products, err error) {
	products, errGet := serv.repository.GetAll()
	if errGet != nil {
		err = errGet
		return
	}
	return
}

func (serv *service) Save(product products) (prod products, err error) {
	lastId, errLastId := serv.repository.GetLastId()
	if errLastId != nil {
		err = errLastId
		return
	}

	prod, errGet := serv.repository.Save()
}
