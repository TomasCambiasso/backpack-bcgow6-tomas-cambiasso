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
	filePtr := file.File{}
	filePtr.SetPath("./tickets.csv")
	tickets, readError := filePtr.Read()
	err := errors.Unwrap(readError)
	for err != nil {
		fmt.Println(err)
		err = errors.Unwrap(err)
	}

	booking := service.NewBookings(tickets)

	ticket := service.Ticket{
		Names:       "Tomas Cambiasso",
		Email:       "algo",
		Destination: "Buenos Aires",
		Price:       10,
		Date:        "12:00",
	}

	// Create //
	booking.Create(ticket)
	// Read //
	ticket, ticketErr := booking.Read(1001)
	if ticketErr != nil {
		fmt.Println(ticketErr)
	}

	fmt.Println(ticket)
	// Read failure //
	fmt.Println("---Read failure---")
	ticket, ticketErr = booking.Read(1002)
	if ticketErr != nil {
		fmt.Println(ticketErr)
	}

	// UPDATE //
	fmt.Println("---Update---")
	ticket = service.Ticket{
		Id:          1001,
		Names:       "Tomas Cambiassoooooo",
		Email:       "algo",
		Destination: "Buenos Aires",
		Price:       10,
		Date:        "12:00",
	}

	ticket, ticketErr = booking.Update(ticket.Id, ticket)
	if ticketErr != nil {
		fmt.Println(ticketErr)
	}

	ticket, ticketErr = booking.Read(1001)
	if ticketErr != nil {
		fmt.Println(ticketErr)
	}

	fmt.Println(ticket)
	// UPDATE FAILURE //
	fmt.Println("---Update failure---")
	ticket = service.Ticket{
		Id:          1002,
		Names:       "Tomas Cambiassoooooo",
		Email:       "algo",
		Destination: "Buenos Aires",
		Price:       10,
		Date:        "12:00",
	}

	ticket, ticketErr = booking.Update(ticket.Id, ticket)
	if ticketErr != nil {
		fmt.Println(ticketErr)
	}
	// DELETE //
	fmt.Println("---Delete---")
	_, ticketErr = booking.Delete(1001)
	if ticketErr != nil {
		fmt.Println(ticketErr)
	}

	ticket, ticketErr = booking.Read(1001)
	if ticketErr != nil {
		fmt.Println(ticketErr)
	}

	// WRITE //
	ticket = service.Ticket{
		Id:          1002,
		Names:       "Tomas Cambiassoooooo",
		Email:       "algo",
		Destination: "Buenos Aires",
		Price:       10,
		Date:        "12:00",
	}
	booking.Create(ticket)
	writeErr := filePtr.Write(booking.GetTickets())
	if writeErr != nil {
		fmt.Println(ticketErr)
	}
}
