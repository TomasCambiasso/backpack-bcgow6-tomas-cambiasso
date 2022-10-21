package transactions

import "github.com/TomasCambiasso/backpack-bcgow6-tomas-cambiasso/internal/domain"

type Service interface {
	GetAll() ([]domain.Transaction, error)
	Store(transaction_code, moneda, emisor, receptor, transaction_date string, monto float64) (domain.Transaction, error)
	Update(id int, transaction_code, moneda, emisor, receptor, transaction_date string, monto float64) (domain.Transaction, error)
	UpdateCodeAndAmount(id int, transaction_code string, monto float64) (domain.Transaction, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]domain.Transaction, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) Store(transaction_code, moneda, emisor, receptor, transaction_date string, monto float64) (domain.Transaction, error) {
	t, err := s.repository.Store(transaction_code, moneda, emisor, receptor, transaction_date, monto)
	if err != nil {
		return domain.Transaction{}, err
	}

	return t, nil
}

func (s *service) Update(id int, transaction_code, moneda, emisor, receptor, transaction_date string, monto float64) (domain.Transaction, error) {
	return s.repository.Update(id, transaction_code, moneda, emisor, receptor, transaction_date, monto)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) UpdateCodeAndAmount(id int, transaction_code string, monto float64) (domain.Transaction, error) {
	return s.repository.UpdateCodeAndAmount(id, transaction_code, monto)
}
