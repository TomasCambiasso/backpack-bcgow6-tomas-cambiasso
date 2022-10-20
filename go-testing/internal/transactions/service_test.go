package transactions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {

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
	service := NewService(repository)

	expectedTrans := transaction{
		Id:               2,
		Transaction_code: "After Update",
		Moneda:           "After Update",
		Monto:            9999,
		Emisor:           "After Update",
		Receptor:         "After Update",
		Transaction_date: "4/10/2022",
	}

	updT, err := service.Update(2, expectedTrans.Transaction_code, expectedTrans.Moneda, expectedTrans.Emisor, expectedTrans.Moneda, expectedTrans.Transaction_date, expectedTrans.Monto)

	assert.Nil(t, err)
	assert.True(t, db.readCheck)
	assert.Equal(t, expectedTrans, updT)
}

func TestDeleteExistentId(t *testing.T) {

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
	service := NewService(repository)

	err := service.Delete(2)
	assert.Nil(t, err)
	err = service.Delete(2)
	assert.Error(t, err)
}

func TestDeleteNonExistentId(t *testing.T) {

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
	service := NewService(repository)

	expectedTrans := transaction{
		Id:               2,
		Transaction_code: "After Update",
		Moneda:           "After Update",
		Monto:            9999,
		Emisor:           "After Update",
		Receptor:         "After Update",
		Transaction_date: "4/10/2022",
	}

	updT, err := service.Update(2, expectedTrans.Transaction_code, expectedTrans.Moneda, expectedTrans.Emisor, expectedTrans.Moneda, expectedTrans.Transaction_date, expectedTrans.Monto)

	assert.Nil(t, err)
	assert.True(t, db.readCheck)
	assert.Equal(t, expectedTrans, updT)
}
