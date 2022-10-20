package main

import (
	"fmt"
	"os"

	"github.com/TomasCambiasso/backpack-bcgow6-tomas-cambiasso/cmd/server/handler"
	"github.com/TomasCambiasso/backpack-bcgow6-tomas-cambiasso/internal/transactions"
	"github.com/TomasCambiasso/backpack-bcgow6-tomas-cambiasso/pkg/store"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	pr := r.Group("/transactions")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.DELETE("/:id", p.Delete())
	pr.PATCH("/:id", p.UpdateCodeAndAmount())
	r.Run()
}
