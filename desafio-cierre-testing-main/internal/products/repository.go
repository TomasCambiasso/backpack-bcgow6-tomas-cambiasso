package products

import "errors"

type Repository interface {
	GetAllBySeller(sellerID string) ([]Product, error)
}

type repository struct {
	db map[string][]Product
}

func NewRepository(db map[string][]Product) Repository {
	return &repository{db: db}
}

func (r *repository) GetAllBySeller(sellerID string) ([]Product, error) {
	products, ok := r.db[sellerID]
	if !ok {
		return nil, errors.New("seller_id not found")
	}
	return products, nil
}
