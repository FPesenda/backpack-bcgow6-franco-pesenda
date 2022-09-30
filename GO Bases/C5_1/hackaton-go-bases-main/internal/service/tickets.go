package service

import (
	"errors"
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
	b.Tickets = append(b.Tickets, t)
	return t, nil
}

func (b *bookings) Read(id int) (tik Ticket, err error) {
	for _, v := range b.Tickets {
		if v.Id == id {
			tik = v
		}
	}
	if &tik == nil {
		err = errors.New("Ticket no encontrado")
	}
	return
}

func (b *bookings) Update(id int, t Ticket) (ttchange Ticket, err error) {
	ttchange, err = b.Read(t.Id)
	if err != nil {
		err = errors.New("No puedo actualizar el ticket porque no existe")
	} else {
		ttchange = Ticket{
			Id:          t.Id,
			Names:       t.Names,
			Email:       t.Email,
			Destination: t.Destination,
			Date:        t.Date,
			Price:       t.Price,
		}
	}
	return
}

func (b *bookings) Delete(id int) (isDeleted int, err error) {
	var indexOfElementToDelete int
	var elementToDelete Ticket
	for i, v := range b.Tickets {
		if v.Id == id {
			indexOfElementToDelete = i
			elementToDelete = v
		}
	}
	if &elementToDelete == nil {
		err = errors.New("No se encontro el elemento a borrar")
		isDeleted = 0
	} else {
		b.Tickets[indexOfElementToDelete] = b.Tickets[len(b.Tickets)-1]
		b.Tickets = b.Tickets[:len(b.Tickets)-1]
		isDeleted = 1
	}
	return
}
