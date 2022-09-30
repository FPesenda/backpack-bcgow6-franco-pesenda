package file

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

type File struct {
	path string
}

func (f *File) Read() (readTickets []service.Ticket, err error) {
	fileOpen, errOpen := os.Open(f.path)
	if errOpen != nil {
		err = errOpen
	}
	reader := csv.NewReader(fileOpen)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic("Error en la lectura del Ticket")
		}
		idTemp, _ := strconv.Atoi(line[0])
		priceTemp, _ := strconv.Atoi(line[5])
		readTickets = append(readTickets, service.Ticket{
			Id:          idTemp,
			Names:       string(line[1]),
			Email:       string(line[2]),
			Destination: string(line[3]),
			Date:        string(line[4]),
			Price:       priceTemp,
		})
	}
	return
}

// TENGO QUE MEJORAR EL NAEJO DE ERRORES
func (f *File) Write(ticket service.Ticket) (err error) {
	fileOpen, errOpen := os.Open(f.path)
	_, errWrite := fileOpen.WriteString(fmt.Sprint(
		ticket.Id, ",",
		ticket.Names, ",",
		ticket.Email, ",",
		ticket.Destination, ",",
		ticket.Date, ",",
		ticket.Price, ",",
	))
	if errWrite != nil {
		err = errWrite
	}
	if errOpen != nil {
		err = errOpen
	}
	return
}
