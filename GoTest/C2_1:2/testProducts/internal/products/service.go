package products

type service struct {
	repository Repository
}

type Service interface {
	Patch(int, string, float64) (Products, error)
	Delete(int) error
	GetAll() ([]Products, error)
	Store(string, string, string, float64) (Products, error)
	UpdateByID(id int, name, color, code string, price float64) (product Products, err error)
}

func NewService(rep Repository) Service {
	return &service{
		repository: rep,
	}
}

func (serv *service) Patch(id int, name string, price float64) (product Products, err error) {
	product, err = serv.repository.Patch(id, name, price)
	return
}

func (serv *service) Delete(id int) (err error) {
	_, err = serv.repository.Delete(id)
	return
}

func (s *service) GetAll() ([]Products, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (serv *service) Store(name, color, code string, price float64) (Products, error) {
	lastID, err := serv.repository.GetLastId()
	if err != nil {
		return Products{}, err
	}

	lastID++

	producto, err := serv.repository.Store(lastID, name, color, code, price)
	if err != nil {
		return Products{}, err
	}

	return producto, nil
}

func (serv *service) UpdateByID(id int, name, color, code string, price float64) (product Products, err error) {
	product, err = serv.repository.UpdateByID(id, name, color, code, price)
	return
}
