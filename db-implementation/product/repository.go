package product

import (
	"context"
	"database/sql"
	"db-implementation/domain"
)

type Repository interface {
	GetByID(ctx context.Context, id int) (domain.Product, error)
	GetAll(ctx context.Context) ([]domain.Product, error)
	Store(ctx context.Context, name string, ptype string, count int, price float64) (int, error)
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

const (
	SAVE_PRODUCT = "INSERT INTO products (name, type, count, price) VALUES (?, ?, ?, ?);"
	GET_ALL_FULL = "SELECT p.id, p.name, p.type, p.count, p.price, w.name, w.adress FROM products p inner join warehouses w on p.id_warehouse = w.id;"
	GET_ALL      = "SELECT p.id, p.name, p.type, p.count, p.price FROM products p"
	GET_PRODUCT  = "SELECT id, name, type, count, price FROM products WHERE id=?;"
	DELETE       = "DELETE FROM products WHERE id = ?"
)

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) Store(ctx context.Context, name string, ptype string, count int, price float64) (int, error) {
	stm, err := r.db.Prepare(SAVE_PRODUCT) //preparamos la consulta
	if err != nil {
		return 0, err
	}
	defer stm.Close()
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

func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	var prods []domain.Product
	rows, err := r.db.Query(GET_ALL)
	if err != nil {
		return prods, err
	}
	for rows.Next() {
		var prod domain.Product
		if err := rows.Scan(&prod.Id, &prod.Name, &prod.Ptype, &prod.Count, &prod.Price); err != nil {
			return prods, err
		}
		prods = append(prods, prod)
	}

	return prods, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {

	stmt, err := r.db.Prepare(DELETE)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id) // retorna un sql.Result y un error

	if err != nil {
		return err
	}

	return nil
}
