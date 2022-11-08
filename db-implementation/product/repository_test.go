package product

import (
	"context"
	"db-implementation/domain"
	"db-implementation/pkg/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {

	_, db := db.ConnectDatabase()
	repo := NewRepository(db)
	id, err := repo.Store(context.TODO(), "ProdB", "TipoB", 1, 10.0)
	assert.Equal(t, 5, id)
	assert.Nil(t, err)

}

func TestGetByID(t *testing.T) {

	_, db := db.ConnectDatabase()
	repo := NewRepository(db)
	actualProd, err := repo.GetByID(context.TODO(), 5)
	expectedProd := domain.Product{
		Id: 5,
		Name: "ProdB",
		Ptype: "TipoB",
		Count: 1,
		Price: 10.0,
	}
	assert.Nil(t, err)
	assert.Equal(t, expectedProd, actualProd)

}

