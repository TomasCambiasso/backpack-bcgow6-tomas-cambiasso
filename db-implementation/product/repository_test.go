package product

import (
	"context"
	"db-implementation/domain"
	"db-implementation/pkg/db"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {

	db, err := db.InitDb()
	assert.NoError(t, err)
	repo := NewRepository(db)
	id, err := repo.Store(context.TODO(), "ProdB", "TipoB", 1, 10.0, 1)
	assert.NoError(t, err)
	prod, err := repo.GetByID(context.TODO(), id)
	assert.NoError(t, err)
	expectedProd := domain.Product{
		Id:    8,
		Name:  "ProdB",
		Ptype: "TipoB",
		Count: 1,
		Price: 10.0,
	}
	assert.Equal(t, expectedProd, prod)

}

func TestUpdate(t *testing.T) {

	db, mock, err := sqlmock.New()
	columns := []string{"id", "name", "type", "count", "price", "id_warehouse"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(1, "ProdA", "ProdB", 1, 10.0, 1)
	mock.ExpectPrepare(regexp.QuoteMeta(UPDATE_PRODUCT))
	mock.ExpectExec(regexp.QuoteMeta(UPDATE_PRODUCT)).WithArgs("ProdC", "TipoC", 2, 20.0, 1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery(regexp.QuoteMeta(GET_PRODUCT)).WillReturnRows(rows)
	assert.NoError(t, mock.ExpectationsWereMet())
	repo := NewRepository(db)
	id, err := repo.Update(context.TODO(), "ProdC", "TipoC", 2, 20.0, 1, 1)
	assert.NoError(t, err)
	prod, err := repo.GetByID(context.TODO(), id)
	assert.NoError(t, err)
	expectedProd := domain.Product{
		Id:    1,
		Name:  "ProdC",
		Ptype: "TipoC",
		Count: 2,
		Price: 20.0,
	}
	assert.Equal(t, expectedProd, prod)

}

func TestGetByID(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)
	repo := NewRepository(db)
	actualProd, err := repo.GetByID(context.TODO(), 4)
	expectedProd := domain.Product{
		Id:    4,
		Name:  "ProdA",
		Ptype: "TipoA",
		Count: 1,
		Price: 10.0,
	}
	assert.Nil(t, err)
	assert.Equal(t, expectedProd, actualProd)

}

func TestGetAll(t *testing.T) {

	_, db := db.ConnectDatabase()
	repo := NewRepository(db)
	actualProds, err := repo.GetAll(context.TODO())
	expectedProds := []domain.Product{
		{
			Id:    4,
			Name:  "ProdA",
			Ptype: "TipoA",
			Count: 1,
			Price: 10.0,
		},
		{
			Id:    5,
			Name:  "ProdB",
			Ptype: "TipoB",
			Count: 1,
			Price: 10.0,
		},
	}
	assert.Nil(t, err)
	assert.Equal(t, expectedProds, actualProds)

}

func TestDelete(t *testing.T) {

	_, db := db.ConnectDatabase()
	repo := NewRepository(db)
	err := repo.Delete(context.TODO(), 5)
	assert.Nil(t, err)

}
