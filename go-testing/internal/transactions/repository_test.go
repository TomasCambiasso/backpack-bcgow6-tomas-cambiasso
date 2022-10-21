package transactions

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubDB struct {
}

func (st StubDB) Read(data interface{}) error {
	transactions := []transaction{
		{
			Id:               2,
			Transaction_code: "000A",
			Moneda:           "EU",
			Monto:            30,
			Emisor:           "Jose Juan",
			Receptor:         "Tomas Cambiasso",
			Transaction_date: "4/10/2022",
		},
		{
			Id:               3,
			Transaction_code: "0010",
			Moneda:           "US",
			Monto:            40,
			Emisor:           "Ladimus Postalo",
			Receptor:         "Jose Juan",
			Transaction_date: "5/10/2022",
		},
	}

	transData := data.(*[]transaction)
	*transData = append(*transData, transactions...)

	return nil
}

func (st StubDB) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {

	db := StubDB{}
	repository := NewRepository(db)

	trans, err := repository.GetAll()
	if err != nil {
		println("algo")
	}

	expectedTrans := []transaction{
		{
			Id:               2,
			Transaction_code: "000A",
			Moneda:           "EU",
			Monto:            30,
			Emisor:           "Jose Juan",
			Receptor:         "Tomas Cambiasso",
			Transaction_date: "4/10/2022",
		},
		{
			Id:               3,
			Transaction_code: "0010",
			Moneda:           "US",
			Monto:            40,
			Emisor:           "Ladimus Postalo",
			Receptor:         "Jose Juan",
			Transaction_date: "5/10/2022",
		},
	}

	assert.Equal(t, expectedTrans, trans)
}

type MockDB struct {
	readCheck    bool
	readError    bool
	writeError   bool
	transactions []transaction
}

func (mk *MockDB) Read(data interface{}) error {
	transData := data.(*[]transaction)
	if mk.readError {
		return errors.New("couldn't read")
	}
	*transData = append(*transData, mk.transactions...)
	mk.readCheck = true
	return nil
}

func (mk *MockDB) Write(data interface{}) error {
	if mk.writeError {
		return errors.New("couldn't write")
	}
	transData := data.(*[]transaction)
	mk.transactions = *transData
	return nil
}

func TestUpdateName(t *testing.T) {
	transactions := []transaction{
		{
			Id:               2,
			Transaction_code: "Before Update",
			Moneda:           "EU",
			Monto:            30,
			Emisor:           "Jose Juan",
			Receptor:         "Tomas Cambiasso",
			Transaction_date: "4/10/2022",
		},
		{
			Id:               3,
			Transaction_code: "0010",
			Moneda:           "US",
			Monto:            40,
			Emisor:           "Ladimus Postalo",
			Receptor:         "Jose Juan",
			Transaction_date: "5/10/2022",
		},
	}

	db := MockDB{transactions: transactions}
	repository := NewRepository(&db)

	expectedTrans := transaction{
		Id:               2,
		Transaction_code: "After Update",
		Moneda:           "EU",
		Monto:            9999,
		Emisor:           "Jose Juan",
		Receptor:         "Tomas Cambiasso",
		Transaction_date: "4/10/2022",
	}

	newT, err := repository.UpdateCodeAndAmount(2, "After Update", 9999)
	if err != nil {
		println("algo")
	}
	assert.True(t, db.readCheck)
	assert.Equal(t, expectedTrans, newT)

}

func TestUpdateNameNotFound(t *testing.T) {

	db := MockDB{}
	repository := NewRepository(&db)
	newT, err := repository.UpdateCodeAndAmount(2, "After Update", 9999)
	assert.ErrorContains(t, err, "no encontrada")
	assert.Equal(t, transaction{}, newT)

}

func TestStoreOk(t *testing.T) {
	transactions := []transaction{
		{
			Id:               2,
			Transaction_code: "Before Update",
			Moneda:           "EU",
			Monto:            30,
			Emisor:           "Jose Juan",
			Receptor:         "Tomas Cambiasso",
			Transaction_date: "4/10/2022",
		},
		{
			Id:               3,
			Transaction_code: "0010",
			Moneda:           "US",
			Monto:            40,
			Emisor:           "Ladimus Postalo",
			Receptor:         "Jose Juan",
			Transaction_date: "5/10/2022",
		},
	}

	db := MockDB{transactions: transactions}
	repository := NewRepository(&db)

	expectedTrans := transaction{
		Id:               4,
		Transaction_code: "3",
		Moneda:           "3",
		Monto:            9999,
		Emisor:           "3",
		Receptor:         "3",
		Transaction_date: "4/10/2022",
	}

	newT, err := repository.Store(expectedTrans.Transaction_code, expectedTrans.Moneda, expectedTrans.Emisor, expectedTrans.Receptor, expectedTrans.Transaction_date, expectedTrans.Monto)
	if err != nil {
		println("algo")
	}
	assert.True(t, db.readCheck)
	assert.Equal(t, expectedTrans, newT)

}

func TestStoreFirstElement(t *testing.T) {
	transactions := []transaction{}

	db := MockDB{transactions: transactions}
	repository := NewRepository(&db)

	expectedTrans := transaction{
		Id:               1,
		Transaction_code: "3",
		Moneda:           "3",
		Monto:            9999,
		Emisor:           "3",
		Receptor:         "3",
		Transaction_date: "4/10/2022",
	}

	newT, err := repository.Store(expectedTrans.Transaction_code, expectedTrans.Moneda, expectedTrans.Emisor, expectedTrans.Receptor, expectedTrans.Transaction_date, expectedTrans.Monto)
	if err != nil {
		println("algo")
	}
	assert.True(t, db.readCheck)
	assert.Equal(t, expectedTrans, newT)

}
