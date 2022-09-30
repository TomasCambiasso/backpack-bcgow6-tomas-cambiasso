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

func (b *bookings) Read(id int) (Ticket, error) {
	defer func() error {
		recover()
		return fmt.Errorf("error: el ticket %d pedido no existe", id)
	}()
	ticket := b.Tickets[id]
	return ticket, nil
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	b.mu.Lock()
	defer func() error {
		recover()
		b.mu.Unlock()
		return fmt.Errorf("error: el ticket %d pedido no existe", id)
	}()
	b.Tickets[id] = t
	return t, nil
}

func (b *bookings) Delete(id int) (int, error) {
	b.mu.Lock()
	defer func() error {
		recover()
		b.mu.Unlock()
		return nil
	}()
	b.Tickets[id] = b.Tickets[len(b.Tickets)-1]
	return 0, nil
}
