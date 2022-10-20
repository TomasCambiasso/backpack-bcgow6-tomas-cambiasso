package transactions

import (
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
	transactions []transaction
}

func (mk *MockDB) Read(data interface{}) error {
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
	mk.transactions = transactions
	transData := data.(*[]transaction)
	*transData = append(*transData, transactions...)
	mk.readCheck = true
	return nil
}

func (mk MockDB) Write(data interface{}) error {

	transData := data.([]transaction)
	mk.transactions = transData
	return nil
}

func TestUpdateName(t *testing.T) {
	db := MockDB{}
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
