package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"
)

type transaction struct {
	Id               int     `json:"id"`
	Transaction_code string  `json:"transaction_code"`
	Moneda           string  `json:"moneda"`
	Monto            float64 `json:"monto"`
	Emisor           string  `json:"emisor"`
	Receptor         string  `json:"receptor"`
	Transaction_date string  `json:"transaction_date"`
}

func filterByField(ctx *gin.Context, t transaction){
	// Funcion que deberia validar que, si los haya, cada ctx.Query(campo_struct)
}

func GetFiltererd(ts []transaction) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var filtered []*transaction
		for _, t := range ts {
			if ctx.Query("moneda") == t.Moneda {
				filtered = append(filtered, &t)
			}
		}
		ctx.JSON(200, filtered)
	}
}

func GetTransaction(ts map[int]transaction) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.String(404, "Not found")
		}
		transaction, ok := ts[id]
		if ok {
			ctx.String(200, "%+v", transaction)
		} else {
			ctx.String(404, "Not found")
		}
	}
}

func main() {

	data, fileError := ioutil.ReadFile("./../users.json")
	if fileError != nil {
		fmt.Println(fileError.Error())
		return
	}
	var transactions []transaction
	err := json.Unmarshal(data, &transactions)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	transMap := make(map[int]transaction)

	for _, t := range transactions {
		transMap[t.Id] = t
	}

	router := gin.Default()
	router.GET("/transacciones", GetFiltererd(transactions))
	router.GET("/transacciones/:id", GetTransaction(transMap))
	router.Run()

}
