package transactions

type Service interface {
	GetAll() ([]transaction, error)
	Store(transaction_code, moneda, emisor, receptor, transaction_date string, monto float64) (transaction, error)
	Update(id int, transaction_code, moneda, emisor, receptor, transaction_date string, monto float64) (transaction, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]transaction, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) Store(transaction_code, moneda, emisor, receptor, transaction_date string, monto float64) (transaction, error) {
	t, err := s.repository.Store(transaction_code, moneda, emisor, receptor, transaction_date, monto)
	if err != nil {
		return transaction{}, err
	}

	return t, nil
}

func (s *service) Update(id int, transaction_code, moneda, emisor, receptor, transaction_date string, monto float64) (transaction, error) {
	return s.repository.Update(id, transaction_code, moneda, emisor, receptor, transaction_date, monto)
}
