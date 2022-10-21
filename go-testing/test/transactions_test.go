package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TomasCambiasso/backpack-bcgow6-tomas-cambiasso/cmd/server/handler"
	"github.com/TomasCambiasso/backpack-bcgow6-tomas-cambiasso/internal/domain"
	"github.com/TomasCambiasso/backpack-bcgow6-tomas-cambiasso/internal/transactions"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(mk MockDB) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	repo := transactions.NewRepository(&mk)
	service := transactions.NewService(repo)
	p := handler.NewTransaction(service)

	r := gin.Default()

	pr := r.Group("/transactions")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func TestUpdateProduct(t *testing.T) {
	// arrrange
	mockDB := MockDB{
		Transactions: []domain.Transaction{
			{
				Id:               2,
				Transaction_code: "000A",
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
		}}
	var resp domain.Transaction
	r := createServer(mockDB)
	req, rr := createRequestTest(http.MethodPut, "/transactions/2", `{
        "Transaction_code": "2","Moneda": "2","Monto": 10,"Emisor": "2","Receptor":"2","Transaction_date":"4/10/2022"
    }`)
	// act
	r.ServeHTTP(rr, req)
	// assert
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
}
