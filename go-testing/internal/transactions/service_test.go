package transactions

import (
	"testing"

	"github.com/TomasCambiasso/backpack-bcgow6-tomas-cambiasso/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestUpdateReadErr(t *testing.T) {

	db := MockDB{}
	db.ReadError = true
	repository := NewRepository(&db)
	service := NewService(repository)

	expectedTrans := domain.Transaction{
		Id:               2,
		Transaction_code: "After Update",
		Moneda:           "After Update",
		Monto:            9999,
		Emisor:           "After Update",
		Receptor:         "After Update",
		Transaction_date: "4/10/2022",
	}

	updT, err := service.Update(2, expectedTrans.Transaction_code, expectedTrans.Moneda, expectedTrans.Emisor, expectedTrans.Moneda, expectedTrans.Transaction_date, expectedTrans.Monto)

	assert.ErrorContains(t, err, "couldn't read")
	assert.Equal(t, domain.Transaction{}, updT)
}

func TestUpdateWriteErr(t *testing.T) {
	transactions := []domain.Transaction{
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

	db := MockDB{Transactions: transactions}
	db.WriteError = true
	repository := NewRepository(&db)
	service := NewService(repository)

	expectedTrans := domain.Transaction{
		Id:               2,
		Transaction_code: "After Update",
		Moneda:           "After Update",
		Monto:            9999,
		Emisor:           "After Update",
		Receptor:         "After Update",
		Transaction_date: "4/10/2022",
	}

	updT, err := service.Update(2, expectedTrans.Transaction_code, expectedTrans.Moneda, expectedTrans.Emisor, expectedTrans.Moneda, expectedTrans.Transaction_date, expectedTrans.Monto)

	assert.ErrorContains(t, err, "couldn't write")
	assert.Equal(t, domain.Transaction{}, updT)
}

func TestUpdateNotFound(t *testing.T) {

	db := MockDB{}
	repository := NewRepository(&db)
	service := NewService(repository)

	expectedTrans := domain.Transaction{
		Id:               2,
		Transaction_code: "After Update",
		Moneda:           "After Update",
		Monto:            9999,
		Emisor:           "After Update",
		Receptor:         "After Update",
		Transaction_date: "4/10/2022",
	}

	updT, err := service.Update(2, expectedTrans.Transaction_code, expectedTrans.Moneda, expectedTrans.Emisor, expectedTrans.Moneda, expectedTrans.Transaction_date, expectedTrans.Monto)

	assert.ErrorContains(t, err, "no encontrada")
	assert.Equal(t, domain.Transaction{}, updT)
}

func TestUpdateOk(t *testing.T) {

	transactions := []domain.Transaction{
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

	db := MockDB{Transactions: transactions}
	repository := NewRepository(&db)
	service := NewService(repository)

	expectedTrans := domain.Transaction{
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
	assert.True(t, db.ReadCheck)
	assert.Equal(t, expectedTrans, updT)
}

func TestDeleteExistentId(t *testing.T) {

	transactions := []domain.Transaction{
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

	db := MockDB{Transactions: transactions}
	repository := NewRepository(&db)
	service := NewService(repository)

	err := service.Delete(2)
	assert.Nil(t, err)
	err = service.Delete(2)
	assert.Error(t, err)
}

func TestDeleteNonExistentId(t *testing.T) {

	transactions := []domain.Transaction{
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

	db := MockDB{Transactions: transactions}
	repository := NewRepository(&db)
	service := NewService(repository)

	expectedTrans := domain.Transaction{
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
	assert.True(t, db.ReadCheck)
	assert.Equal(t, expectedTrans, updT)
}
