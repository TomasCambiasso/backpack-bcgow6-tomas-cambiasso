package main

import (
	"errors"
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func main() {
	var tickets []service.Ticket
	// Funcion para obtener tickets del archivo csv
	service.NewBookings(tickets)
	filePtr := file.File{}
	filePtr.SetPath("./tickets.csv")
	tickets, readError := filePtr.Read()
	err := errors.Unwrap(readError)
	for err != nil{
		fmt.Println(err)
		err = errors.Unwrap(err)
	}
	
}
