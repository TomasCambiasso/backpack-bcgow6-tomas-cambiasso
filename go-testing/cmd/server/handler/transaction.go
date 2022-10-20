package handler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/TomasCambiasso/backpack-bcgow6-tomas-cambiasso/internal/transactions"
	"github.com/TomasCambiasso/backpack-bcgow6-tomas-cambiasso/pkg/web"

	"github.com/gin-gonic/gin"
)

type request struct {
	Transaction_code string   `json:"transaction_code"`
	Moneda           string   `json:"moneda"`
	Monto            *float64 `json:"monto"`
	Emisor           string   `json:"emisor"`
	Receptor         string   `json:"receptor" `
	Transaction_date string   `json:"transaction_date"`
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(s transactions.Service) *Transaction {
	return &Transaction{
		service: s,
	}
}

// ListTransactions godoc
// @Summary List transactions
// @Tags Transactions
// @Description get transactions
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /transactions [get]
func (t *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token inválido"))
			return
		}

		ts, err := t.service.GetAll()
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}
		if len(ts) == 0 {
			ctx.JSON(404, web.NewResponse(404, nil, "No hay productos almacenados"))
			return

		}
		ctx.JSON(200, web.NewResponse(200, ts, ""))
	}
}
// StoreTransactions godoc
// @Summary Store transactions
// @Tags Transactions
// @Description store transactions
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Transaction to store"
// @Success 200 {object} web.Response
// @Router /transactions [post]
func (t *Transaction) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token inválido"))
			return
		}
		var req request
		allGood := validateFields(ctx, &req)
		if !allGood {
			return
		}
		nt, err := t.service.Store(req.Transaction_code, req.Moneda, req.Emisor, req.Receptor, req.Transaction_date, *req.Monto)
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, nt, ""))
	}
}

// UpdateTransactions godoc
// @Summary Update transactions
// @Tags Transactions
// @Description update transactions, providing all values
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "Transaction ID"
// @Param product body request true "All transaction values to update"
// @Success 200 {object} web.Response
// @Router /transactions [put]
func (t *Transaction) Update() gin.HandlerFunc { /// Las validaciones deberian ser una funcion aparte dado que tambien se deberian usar en store
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token inválido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "ID invalida"))
			return
		}
		var req request
		allGood := validateFields(ctx, &req)
		if !allGood {
			return
		}
		nt, err := t.service.Update(int(id), req.Transaction_code, req.Moneda, req.Emisor, req.Receptor, req.Transaction_date, *req.Monto)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, nt, ""))
	}
}
// DeleteTransaction godoc
// @Summary Delete transaction
// @Tags Transactions
// @Description Delete transactions, providing id to delete
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "Transaction ID"
// @Success 200 {object} web.Response
// @Router /transactions{id} [delete]
func (t *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "ID invalida"))
			return
		}
		err = t.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		response := fmt.Sprintf("La transaccion %d ha sido eliminada", id)
		ctx.JSON(200, web.NewResponse(200, response, ""))
	}
}

// UpdateTransactionsCodeAndAmount godoc
// @Summary UpdateCodeAndAmount transactions
// @Tags Transactions
// @Description update transaction code and amount
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "Transaction ID to update"
// @Param product body request true "Code and amount to update"
// @Success 200 {object} web.Response
// @Router /transactions [patch]
func (t *Transaction) UpdateCodeAndAmount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "ID invalida"))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.Transaction_code == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El codigo de la transaccion es requerido"))
			return
		}
		if req.Monto == nil {
			ctx.JSON(400, web.NewResponse(400, nil, "El monto de la transaccion es requerido"))
			return
		}
		nt, err := t.service.UpdateCodeAndAmount(int(id), req.Transaction_code, *req.Monto)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, nt)
	}
}

func validateFields(ctx *gin.Context, req *request) bool {
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		return false
	}
	if req.Transaction_code == "" {
		ctx.JSON(400, web.NewResponse(400, nil, "El codigo de la transaccion es requerido"))
		return false
	}
	if req.Moneda == "" {
		ctx.JSON(400, web.NewResponse(400, nil, "La moneda de la transaccion es requerida"))
		return false
	}
	if req.Emisor == "" {
		ctx.JSON(400, web.NewResponse(400, nil, "El emisor es requerido"))
		return false
	}
	if req.Transaction_date == "" {
		ctx.JSON(400, web.NewResponse(400, nil, "La fecha de transaccion es requerida"))
		return false
	}
	if req.Receptor == "" {
		ctx.JSON(400, web.NewResponse(400, nil, "El receptor es requerido"))
		return false
	}
	if req.Monto == nil {
		ctx.JSON(400, web.NewResponse(400, nil, "El monto de la transaccion es requerido"))
		return false
	}
	return true
}
