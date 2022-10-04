package handler

import (
	"backpack-bcgow6-tomas-cambiasso/C2-TT/internal/transactions"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	Transaction_code string  `json:"transaction_code" binding:"required"`
	Moneda           string  `json:"moneda" binding:"required"`
	Monto            float64 `json:"monto" binding:"required"`
	Emisor           string  `json:"emisor" binding:"required"`
	Receptor         string  `json:"receptor" binding:"required"`
	Transaction_date string  `json:"transaction_date" binding:"required"`
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

func (c *Transaction) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		p, err := c.service.Store(req.Transaction_code, req.Moneda, req.Emisor, req.Receptor, req.Transaction_date, req.Monto)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *Transaction) Update() gin.HandlerFunc {
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
			ctx.JSON(400, gin.H{"error": "El nombre del producto es requerido"})
			return
		}
		if req.Moneda == "" {
			ctx.JSON(400, gin.H{"error": "El tipo del producto es requerido"})
			return
		}
		if req.Emisor == "" {
			ctx.JSON(400, gin.H{"error": "La cantidad es requerida"})
			return
		}
		if req.Transaction_date == "" {
			ctx.JSON(400, gin.H{"error": "La cantidad es requerida"})
			return
		}
		if req.Receptor == "" {
			ctx.JSON(400, gin.H{"error": "La cantidad es requerida"})
			return
		}
		if req.Monto == 0 {
			ctx.JSON(400, gin.H{"error": "El precio es requerido"})
			return
		}
		t, err := c.service.Update(int(id), req.Transaction_code, req.Moneda, req.Emisor, req.Receptor, req.Transaction_date, req.Monto)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, t)
	}
}
