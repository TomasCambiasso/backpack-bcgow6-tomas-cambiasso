package main

import (
	"fmt"
	"os/exec"
	"time"
	"github.com/gin-gonic/gin"
)

type transactions struct {
	Id               int
	Transaction_code string
	Moneda           string
	Monto            float64
	Emisor           string
	Receptor         string
	Transaction_date string
}

func GetAll(v interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, v)
	}
}

func main() {

	transaction := transactions{
		Id:               3,
		Transaction_code: "001A",
		Moneda:           "US",
		Monto:            10.2,
		Emisor:           "Pepito",
		Receptor:         "Tomas Cambiasso",
		Transaction_date: time.Now().String(),
	}

	router := gin.Default()
	router2 := gin.Default()

	router.GET("/test", func(cxt *gin.Context) {
		cxt.JSON(200, gin.H{
			"message": "Hola Tomas",
		})
	})

	router2.GET("/transacciones", GetAll(transaction))

	go router.Run()
	go router2.Run(":8081")

	cmd := exec.Command("curl", "localhost:8081/transacciones")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(out))

	cmd = exec.Command("curl", "localhost:8080/test")
	out, err = cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(out))

	for {

	}

}
