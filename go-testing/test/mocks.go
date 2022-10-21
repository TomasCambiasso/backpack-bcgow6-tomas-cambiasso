package test

import (
	"errors"

	"github.com/TomasCambiasso/backpack-bcgow6-tomas-cambiasso/internal/transactions"
)


func (mk transactions.m) Read(data interface{}) error {
	transData := data.(*[]transactions.transaction)
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
	transData := data.(*[]transaction)
	mk.Transactions = *transData
	return nil
}
