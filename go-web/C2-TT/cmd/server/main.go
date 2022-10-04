package main

import (
	"backpack-bcgow6-tomas-cambiasso/C2-TT/cmd/server/handler"
	"backpack-bcgow6-tomas-cambiasso/C2-TT/internal/transactions"
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
	r.Run()
}
