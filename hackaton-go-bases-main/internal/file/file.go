package file

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

type File struct {
	path string
}

func (f *File) SetPath(path string) {
	f.path = path
	return
}

func (f *File) Read() ([]service.Ticket, error) {

	filePointer, fileError := os.Open(f.path)
	if fileError != nil {
		return nil, fileError
	}
	defer func() {
		filePointer.Close()
	}()
	records, readError := csv.NewReader(filePointer).ReadAll()
	if fileError != nil {
		return nil, readError
	}
	var tickets []service.Ticket
	var megaError error
	for i, record := range records {
		id, idErr := strconv.Atoi(record[0])
		price, priceErr := strconv.Atoi(record[5])
		if idErr != nil || priceErr != nil {
			megaError = fmt.Errorf("error: precio o id invalidos en record %d, %w", i, megaError)
		}
		ticket := service.Ticket{
			Id:          id,
			Names:       record[1],
			Email:       record[2],
			Destination: record[3],
			Date:        record[4],
			Price:       price,
		}
		tickets = append(tickets, ticket)
	}
	if megaError != nil {

		return nil, megaError
	}

	return tickets, nil
}

func (f *File) Write(tickets []service.Ticket) error {
	filePtr, fileError := os.Create("tickets_new.csv")
	if fileError != nil {
		return fileError
	}
	defer func() {
		filePtr.Close()
	}()
	for _ , ticket := range tickets{
		filePtr.WriteString(ticket.Print()+"\n")
	}
	return nil
}
