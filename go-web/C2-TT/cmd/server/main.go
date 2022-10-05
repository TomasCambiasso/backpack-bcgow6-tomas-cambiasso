package main

import (
	"github.com/TomasCambiasso/backpack-bcgow6-tomas-cambiasso/C2-TT/cmd/server/handler"
	"github.com/TomasCambiasso/backpack-bcgow6-tomas-cambiasso/C2-TT/internal/transactions"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := transactions.NewRepository()
	service := transactions.NewService(repo)

	p := handler.NewTransaction(service)

	r := gin.Default()

	pr := r.Group("/transactions")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.DELETE("/:id", p.Delete())
	pr.PATCH("/:id", p.UpdateCodeAndAmount())
	r.Run()
}
