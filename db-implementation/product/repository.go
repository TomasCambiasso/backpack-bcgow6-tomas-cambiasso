package product

import (
	"context"
	"database/sql"
	"db-implementation/domain"
)

type Repository interface {
	GetByID(ctx context.Context, id int) (domain.Product, error)
	Store(ctx context.Context, name string, ptype string, count int, price float64) (int, error)
}

type repository struct {
	db *sql.DB
}

const (
	SAVE_PRODUCT = "INSERT INTO products (name, type, count, price) VALUES (?, ?, ?, ?);"

	GET_PRODUCT = "SELECT id, name, type, count, price FROM products WHERE id=?;"
)

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) Store(ctx context.Context, name string, ptype string, count int, price float64) (int, error) {
	stm, err := r.db.Prepare(SAVE_PRODUCT) //preparamos la consulta
	if err != nil {
		return 0, err
	}
	//ejecutamos la consulta con aquellos valores a remplazar en la sentencia
	result, err := stm.Exec(name, ptype, count, price)
	if err != nil {
		return 0, err
	}

	//obtenemos el ultimo id
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) GetByID(ctx context.Context, id int) (domain.Product, error) {
	row := r.db.QueryRow(GET_PRODUCT, id)
	var prod domain.Product
	if err := row.Scan(&prod.Id, &prod.Name, &prod.Ptype, &prod.Count, &prod.Price); err != nil {
		return domain.Product{}, err
	}
	return prod, nil
}
