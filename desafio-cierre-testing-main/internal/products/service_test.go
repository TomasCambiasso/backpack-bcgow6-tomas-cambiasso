package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllBySellerOK(t *testing.T) {

	db := make(map[string][]Product)
	db["FEX112AC"] = append(db["FEX112AC"], Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	})
	repository := NewRepository(db)
	service := NewService(repository)

	var expProducts []Product
	expProducts = append(expProducts, Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	})
	actualProducts, err := service.GetAllBySeller("FEX112AC")

	assert.Nil(t, err)
	assert.Equal(t, expProducts, actualProducts)
}
