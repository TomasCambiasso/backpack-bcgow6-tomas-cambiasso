package product

import (
	"context"
	"db-implementation/domain"
	"db-implementation/pkg/db"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var (
	ERROR_FORZADO error = errors.New("ERROR_FORZADO")
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
	newP := domain.Product{
		Id:           1,
		Name:         "ProdC",
		Ptype:        "TipoC",
		Count:        2,
		Price:        20.0,
		Id_warehouse: 1,
	}
	mock.ExpectPrepare(regexp.QuoteMeta(UPDATE_PRODUCT))
	// Le tengo que pasar si o si 6 arguments porque UPDATE_PRODUCT espera 6
	mock.ExpectExec(regexp.QuoteMeta(UPDATE_PRODUCT)).WithArgs(newP.Name, newP.Ptype, newP.Count, newP.Price, newP.Id_warehouse, newP.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	repo := NewRepository(db)
	id, err := repo.Update(context.TODO(), newP.Name, newP.Ptype, newP.Count, newP.Price, newP.Id_warehouse, newP.Id)
	assert.NoError(t, err)
	assert.Equal(t, id, newP.Id)
	assert.NoError(t, mock.ExpectationsWereMet())

}

func TestDeleteSqlMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	repo := NewRepository(db)
	delId := 1
	mock.ExpectPrepare(regexp.QuoteMeta(DELETE))
	mock.ExpectExec(regexp.QuoteMeta(DELETE)).WithArgs(delId).WillReturnResult(sqlmock.NewResult(1, 1))
	err = repo.Delete(context.TODO(), delId)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestDeleteSqlMockError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	repo := NewRepository(db)
	delId := 1
	mock.ExpectPrepare(regexp.QuoteMeta(DELETE))
	mock.ExpectExec(regexp.QuoteMeta(DELETE)).WithArgs(delId).WillReturnError(ERROR_FORZADO)
	err = repo.Delete(context.TODO(), delId)
	assert.EqualValues(t, ERROR_FORZADO, err)
	assert.NoError(t, mock.ExpectationsWereMet())

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
