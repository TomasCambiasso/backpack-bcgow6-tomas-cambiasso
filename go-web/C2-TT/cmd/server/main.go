package main

import (
	"github.com/TomasCambiasso/backpack-bcgow6-tomas-cambiasso/C2-TT/cmd/server/handler"
	"github.com/TomasCambiasso/backpack-bcgow6-tomas-cambiasso/C2-TT/internal/transactions"
	"github.com/TomasCambiasso/backpack-bcgow6-tomas-cambiasso/C2-TT/pkg/store"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("env.env")
	if err != nil {
		fmt.Println(fmt.Errorf("error al intentar cargar archivo .env"))
	}
	db := store.New(store.FileType, "./users.json")
	repo := transactions.NewRepository(db)
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
