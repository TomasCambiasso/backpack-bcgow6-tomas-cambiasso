package service

import (
	"fmt"
	"sync"
)

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
	// Print tickets
	PrintTickets()
	// Get Tickets
	GetTickets() (tickets []Ticket)
}

type bookings struct {
	Tickets []Ticket
	mu      sync.Mutex
}

type Ticket struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if len(b.Tickets) == 0 {
		t.Id = 1
	} else {
		t.Id = b.Tickets[len(b.Tickets)-1].Id + 1
	}
	b.Tickets = append(b.Tickets, t)
	return t, nil
}

func (b *bookings) Read(id int) (returnTicket Ticket, err error) {
	defer func() {
		panic := recover()
		if panic != nil {
			err = fmt.Errorf("error: el ticket %d pedido no existe", id)
		}
	}()
	ticket := b.Tickets[id-1]
	return ticket, nil
}

func (b *bookings) Update(id int, t Ticket) (returnTicket Ticket, err error) {
	b.mu.Lock()
	defer func() {
		panic := recover()
		b.mu.Unlock()
		if panic != nil {
			err = fmt.Errorf("error: el ticket %d pedido no existe", id)
		}
	}()
	b.Tickets[id-1] = t
	return t, nil
}

func (b *bookings) Delete(id int) (int, error) {
	b.mu.Lock()
	defer func() {
		recover()
		b.mu.Unlock()
	}()
	b.Tickets[id-1] = b.Tickets[len(b.Tickets)-1]
	b.Tickets = b.Tickets[:len(b.Tickets)-1]
	// append(slice[: x], slice[x+1]...)
	return 0, nil
}

func (b *bookings) PrintTickets() {
	for _, ticket := range b.Tickets {
		fmt.Println(ticket)
	}
	return
}

func (b *bookings) GetTickets() (tickets []Ticket) {
	return b.Tickets
}

func (t *Ticket) Print() string {
	return fmt.Sprintf("%d,%s,%s,%s,%s,%d", t.Id, t.Names, t.Email, t.Destination, t.Date, t.Price)
}
