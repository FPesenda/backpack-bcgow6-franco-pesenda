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
	pathName := file.File{}
	pathName.Path = "./tickets.csv"
	tickets, err := pathName.Read()
	if err != nil {
		err = errors.New("No se pudo acceder a los elementos")
	}
	booking := service.NewBookings(tickets)

	//MOSTRAR LOS TICKETS ASIGNADOS PARA TRABAJAR EN MEMORIA
	//fmt.Println(booking)
	//CREAR UN NUEVO TICKET
	ticketToWrite, _ := booking.Create(service.Ticket{
		Id:          1001,
		Names:       "Franco Pesenda",
		Email:       "myemail@gmail.com",
		Destination: "USA",
		Date:        "19/11/1995",
		Price:       2500,
	},
	)
	//MOSTRAR EL TICKET 1 LEIDO EN MEMORIA
	fmt.Print("Primer elemento leido ")
	fmt.Println(booking.Read(1))
	//MOSTRAR EL NUEVO TICKET CREADO EN MEMORIA CON EL METODO READ
	fmt.Print("Nuevo Ticket ")
	fmt.Println(booking.Read(1001))
	ticketToWrite, _ = booking.Read(10)
	errWrite := pathName.Write(ticketToWrite)
	if errWrite != nil {
		fmt.Println(errWrite)
	}
}
