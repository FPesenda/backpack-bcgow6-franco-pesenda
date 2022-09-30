package main

import "github.com/bootcamp-go/hackaton-go-bases/internal/service"

//import "github.com/bootcamp-go/hackaton-go-bases/internal/file"

func main() {
	var tickets []service.Ticket
	// Funcion para obtener tickets del archivo csv
	booking := service.NewBookings(tickets)
	booking.Delete(3)
}
