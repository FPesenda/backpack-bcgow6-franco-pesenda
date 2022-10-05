package products

type service struct {
	repository Repository
}

type Service interface {
	GetAll() ([]products, error)
	Save(products) (products, error)
}

func NewService(rep Repository) Service {
	return &service{
		repository: rep,
	}
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
	product.Id = (lastId + 1)

	prod, errGet := serv.repository.Save(product)

	if errGet != nil {
		err = errLastId
		return
	}
	return
}
