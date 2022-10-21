package transactions

import (
	"fmt"

	"github.com/TomasCambiasso/backpack-bcgow6-tomas-cambiasso/internal/domain"
	"github.com/TomasCambiasso/backpack-bcgow6-tomas-cambiasso/pkg/store"
)

var lastID int

// ***Importado**//
type Repository interface {
	GetAll() ([]domain.Transaction, error)
	Store(transaction_code, moneda, emisor, receptor, transaction_date string, monto float64) (domain.Transaction, error)
	LastID() (int, error)
	Update(id int, transaction_code, moneda, emisor, receptor, transaction_date string, monto float64) (domain.Transaction, error)
	UpdateCodeAndAmount(id int, transaction_code string, monto float64) (domain.Transaction, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
} //struct implementa los metodos de la interfaz

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

func (r *repository) Store(transaction_code, moneda, emisor, receptor, transaction_date string, monto float64) (domain.Transaction, error) {
	t := domain.Transaction{
		Transaction_code: transaction_code,
		Moneda:           moneda,
		Monto:            monto,
		Emisor:           emisor,
		Receptor:         receptor,
		Transaction_date: transaction_date,
	}
	var transactions []domain.Transaction
	err := r.db.Read(&transactions)
	if err != nil {
		return domain.Transaction{}, err
	}
	var lastId int

	if len(transactions) == 0 {
		lastId = 0
	} else {
		lastId = transactions[len(transactions)-1].Id
	}
	lastId++
	t.Id = lastId

	transactions = append(transactions, t)

	if err := r.db.Write(&transactions); err != nil {
		return domain.Transaction{}, err
	}

	return t, nil
}

func (r *repository) GetAll() ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	err := r.db.Read(&transactions)
	return transactions, err
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Update(id int, transaction_code, moneda, emisor, receptor, transaction_date string, monto float64) (domain.Transaction, error) {
	t := domain.Transaction{
		Transaction_code: transaction_code,
		Moneda:           moneda,
		Monto:            monto,
		Emisor:           emisor,
		Receptor:         receptor,
		Transaction_date: transaction_date,
	}

	var transactions []domain.Transaction
	err := r.db.Read(&transactions)
	if err != nil {
		return domain.Transaction{}, err
	}

	updated := false
	for i := range transactions {
		if transactions[i].Id == id {
			t.Id = id
			transactions[i] = t
			updated = true
			break
		}
	}
	if !updated {
		return domain.Transaction{}, fmt.Errorf("Transaccion %d no encontrada", id)
	}
	if err := r.db.Write(&transactions); err != nil {
		return domain.Transaction{}, err
	}
	return t, nil
}

func (r *repository) Delete(id int) error {

	var transactions []domain.Transaction
	err := r.db.Read(&transactions)
	if err != nil {
		return err
	}
	deleted := false
	index := -1
	for i := range transactions {
		if transactions[i].Id == id {
			index = i
			deleted = true
			break
		}
	}
	if !deleted {
		return fmt.Errorf("Transaccion %d no encontrada", id)
	}
	transactions = append(transactions[:index], transactions[index+1:]...)

	if err := r.db.Write(&transactions); err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateCodeAndAmount(id int, transaction_code string, monto float64) (domain.Transaction, error) {

	var transactions []domain.Transaction
	err := r.db.Read(&transactions)
	if err != nil {
		return domain.Transaction{}, err
	}

	updated := false
	var t domain.Transaction
	for i := range transactions {
		if transactions[i].Id == id {
			transactions[i].Transaction_code = transaction_code
			transactions[i].Monto = monto
			t = transactions[i]
			updated = true
			break
		}
	}
	if !updated {
		return domain.Transaction{}, fmt.Errorf("Transaccion %d no encontrada", id)
	}

	if err := r.db.Write(&transactions); err != nil {
		return domain.Transaction{}, err
	}
	return t, nil
}
