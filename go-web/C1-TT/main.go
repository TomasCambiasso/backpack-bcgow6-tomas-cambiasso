package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type transaction struct {
	id               int
	Transaction_code string  `json:"transaction_code" binding:"required"`
	Moneda           string  `json:"moneda" binding:"required"`
	Monto            float64 `json:"monto" binding:"required"`
	Emisor           string  `json:"emisor" binding:"required"`
	Receptor         string  `json:"receptor" binding:"required"`
	Transaction_date string  `json:"transaction_date" binding:"required"`
}

func filterByField(ctx *gin.Context, t transaction) {
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

func AddTransaction(ts []transaction) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateHeader(ctx) {
			return
		}
		var t transaction
		if err := ctx.ShouldBindJSON(&t); err != nil {
			var verr validator.ValidationErrors
			if errors.As(err, &verr) {
				var errors []string
				for _, valErr := range verr {
					erro := fmt.Sprintf("El campo %s es requerido", (valErr.Field()))
					errors = append(errors, erro)
				}
				ctx.JSON(400, gin.H{
					"error": errors,
				})
				return
			} else {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err,
				})
				return
			}
		}
		id := len(ts)
		t.id = id
		ts = append(ts, t)
		ctx.JSON(200, t)
		return
	}
}

func validateHeader(ctx *gin.Context) bool {
	token := ctx.GetHeader("token")
	if token != "holi" {
		ctx.JSON(401, gin.H{
			"error": "No tiene permisos para la peticion solicitada",
		})
		return false
	}
	return true
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
		transMap[t.id] = t
	}

	router := gin.Default()
	router.GET("/transacciones", GetFiltererd(transactions))
	router.GET("/transacciones/:id", GetTransaction(transMap))
	router.POST("/transacciones/add", AddTransaction(transactions))
	router.Run()

}
