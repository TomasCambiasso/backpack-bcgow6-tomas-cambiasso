package transactions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubDB struct{

}

func (st StubDB) Read (data interface{}) error{
	transactions := []transaction{
		{
			Id : 2,
			Transaction_code: "000A",
			Moneda: "EU",
			Monto: 30,
			Emisor: "Jose Juan",
			Receptor: "Tomas Cambiasso",
			Transaction_date: "4/10/2022",
		},
		{
			Id : 3,
			Transaction_code: "0010",
			Moneda: "US",
			Monto: 40,
			Emisor: "Ladimus Postalo",
			Receptor: "Jose Juan",
			Transaction_date: "5/10/2022",
		},
	}

	data = transactions
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
			Id : 2,
			Transaction_code: "000A",
			Moneda: "EU",
			Monto: 30,
			Emisor: "Jose Juan",
			Receptor: "Tomas Cambiasso",
			Transaction_date: "4/10/2022",
		},
		{
			Id : 3,
			Transaction_code: "0010",
			Moneda: "US",
			Monto: 40,
			Emisor: "Ladimus Postalo",
			Receptor: "Jose Juan",
			Transaction_date: "5/10/2022",
		},
	}

	assert.Equal(t,expectedTrans,trans)
}