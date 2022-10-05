package transactions

import (
	"fmt"
)

type transaction struct {
	id               int
	Transaction_code string  `json:"transaction_code" binding:"required"`
	Moneda           string  `json:"moneda" binding:"required"`
	Monto            float64 `json:"monto" binding:"required"`
	Emisor           string  `json:"emisor" binding:"required"`
	Receptor         string  `json:"receptor" binding:"required"`
	Transaction_date string  `json:"transaction_date" binding:"required"`
}

var transactions []transaction
var lastID int

// ***Importado**//
type Repository interface {
	GetAll() ([]transaction, error)
	Store(transaction_code, moneda, emisor, receptor, transaction_date string, monto float64) (transaction, error)
	LastID() (int, error)
	Update(id int, transaction_code, moneda, emisor, receptor, transaction_date string, monto float64) (transaction, error)
	UpdateCodeAndAmount(id int, transaction_code string, monto float64) (transaction, error)
	Delete(id int) error
}

type repository struct{} //struct implementa los metodos de la interfaz

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Store(transaction_code, moneda, emisor, receptor, transaction_date string, monto float64) (transaction, error) {
	t := transaction{
		Transaction_code: transaction_code,
		Moneda:           moneda,
		Monto:            monto,
		Emisor:           emisor,
		Receptor:         receptor,
		Transaction_date: transaction_date,
	}
	lastID++
	t.id = lastID
	transactions = append(transactions, t)
	return t, nil
}

func (r *repository) GetAll() ([]transaction, error) {
	return transactions, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Update(id int, transaction_code, moneda, emisor, receptor, transaction_date string, monto float64) (transaction, error) {
	t := transaction{
		Transaction_code: transaction_code,
		Moneda:           moneda,
		Monto:            monto,
		Emisor:           emisor,
		Receptor:         receptor,
		Transaction_date: transaction_date,
	}
	updated := false
	for i := range transactions {
		if transactions[i].id == id {
			t.id = id
			transactions[i] = t
			updated = true
			break
		}
	}
	if !updated {
		return transaction{}, fmt.Errorf("Transaccion %d no encontrada", id)
	}
	return t, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	index := -1
	for i := range transactions {
		if transactions[i].id == id {
			index = i
			deleted = true
			break
		}
	}
	if !deleted {
		return fmt.Errorf("Transaccion %d no encontrada", id)
	}
	transactions = append(transactions[:index], transactions[index+1:]...)
	return nil
}

func (r *repository) UpdateCodeAndAmount(id int, transaction_code string, monto float64) (transaction, error) {
	updated := false
	var t transaction
	for i := range transactions {
		if transactions[i].id == id {
			transactions[i].Transaction_code = transaction_code
			transactions[i].Monto = monto
			t = transactions[i]
			updated = true
			break
		}
	}
	if !updated {
		return transaction{}, fmt.Errorf("Transaccion %d no encontrada", id)
	}
	return t, nil
}
