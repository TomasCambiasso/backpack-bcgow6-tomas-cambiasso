package products

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(db map[string][]Product) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	MapRoutes(r, db)
	return r
}

func MapRoutes(r *gin.Engine, db map[string][]Product) {
	rg := r.Group("/api/v1")
	{
		buildProductsRoutes(rg, db)
	}

}

func buildProductsRoutes(r *gin.RouterGroup, db map[string][]Product) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	prodRoute := r.Group("/products")
	{
		prodRoute.GET("", handler.GetProducts)
	}

}

func TestGetProductsOk(t *testing.T) {

	db := make(map[string][]Product)
	db["FEX112AC"] = append(db["FEX112AC"], Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	})
	r := createServer(db)
	
	req := httptest.NewRequest(http.MethodGet, "/products?seller_id=FEX112AC", nil)
	req.Header.Add("Content-Type", "application/json")
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	var resProducts []Product
	err := json.Unmarshal(res.Body.Bytes(), resProducts)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.Code)

}
