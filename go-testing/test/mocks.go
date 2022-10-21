package test

import (
	"errors"

	"github.com/TomasCambiasso/backpack-bcgow6-tomas-cambiasso/internal/domain"
)
type MockDB struct {
	ReadCheck    bool
	ReadError    bool
	WriteError   bool
	Transactions []domain.Transaction
}


func (mk *MockDB) Read(data interface{}) error {
	transData := data.(*[]domain.Transaction)
	if mk.ReadError {
		return errors.New("couldn't read")
	}
	*transData = append(*transData, mk.Transactions...)
	mk.ReadCheck = true
	return nil
}

func (mk *MockDB) Write(data interface{}) error {
	if mk.WriteError {
		return errors.New("couldn't write")
	}
	transData := data.(*[]domain.Transaction)
	mk.Transactions = *transData
	return nil
}
