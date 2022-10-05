package handler

import (
	"fmt"
	"strconv"

	"github.com/TomasCambiasso/backpack-bcgow6-tomas-cambiasso/C2-TT/internal/transactions"
	"github.com/gin-gonic/gin"
)

type request struct {
	Transaction_code string `json:"transaction_code"`
	Moneda           string `json:"moneda"`
	Monto            string `json:"monto"`
	Emisor           string `json:"emisor"`
	Receptor         string `json:"receptor" `
	Transaction_date string `json:"transaction_date"`
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(s transactions.Service) *Transaction {
	return &Transaction{
		service: s,
	}
}

func (t *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		ts, err := t.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, ts)
	}
}

func (t *Transaction) Store() gin.HandlerFunc { /// Faltan las validaciones see Update()
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		monto, err := strconv.ParseFloat(req.Monto, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "El monto es requerido y es invalido"})
			return
		}
		nt, err := t.service.Store(req.Transaction_code, req.Moneda, req.Emisor, req.Receptor, req.Transaction_date, monto)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, nt)
	}
}

func (t *Transaction) Update() gin.HandlerFunc { /// Las validaciones deberian ser una funcion aparte dado que tambien se deberian usar en store
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.Transaction_code == "" {
			ctx.JSON(400, gin.H{"error": "El codigo de la transaccion es requerido"})
			return
		}
		if req.Moneda == "" {
			ctx.JSON(400, gin.H{"error": "La moneda de la transaccion es requerida"})
			return
		}
		if req.Emisor == "" {
			ctx.JSON(400, gin.H{"error": "El emisor es requerido"})
			return
		}
		if req.Transaction_date == "" {
			ctx.JSON(400, gin.H{"error": "La fecha de transaccion es requerida"})
			return
		}
		if req.Receptor == "" {
			ctx.JSON(400, gin.H{"error": "El receptor es requerido"})
			return
		}
		monto, err := strconv.ParseFloat(req.Monto, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "El monto es requerido y es invalido"})
			return
		}
		nt, err := t.service.Update(int(id), req.Transaction_code, req.Moneda, req.Emisor, req.Receptor, req.Transaction_date, monto)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, nt)
	}
}

func (t *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		err = t.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"data": fmt.Sprintf("La transaccion %d ha sido eliminada", id)})
	}
}

func (t *Transaction) UpdateCodeAndAmount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.Transaction_code == "" {
			ctx.JSON(400, gin.H{"error": "El nombre del producto es requerido"})
			return
		}
		monto, err := strconv.ParseFloat(req.Monto, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "El precio es requerido y es invalido"})
			return
		}
		nt, err := t.service.UpdateCodeAndAmount(id, req.Transaction_code, monto)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, nt)
	}
}
