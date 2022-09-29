package main

import (
	"fmt"
	"os"
)

type Client struct {
	File            int
	NameAndLastName string
	DNI             int
	CellPhone       int
	Address         string
}

var totalClient = 0

func GenerateFileNumber(previous int) (number int) {
	number = (previous + 1)
	totalClient++
	return
}

func NewClient(dni, cp int, nameLastName, addss string) (cl Client, err error) {
	cl = Client{File: GenerateFileNumber(totalClient), NameAndLastName: nameLastName, DNI: dni, CellPhone: cp, Address: addss}
	return
}

func validateClient(client Client) (clientReturn Client, err error) {
	recover()
	if client.File == 0 || client.DNI == 0 || client.CellPhone == 0 {
		err = fmt.Errorf("Cliente con valores nulos")
	} else {
		clientReturn = client
	}
	return
}

func main() {
	client1, err := NewClient(23456765, 0, "Parilli", "DondeVive")
	client2, err := NewClient(1116765, 3514564111, "josua", "DondePerdioElPoncho 1234")
	file, _ := os.Create("customers.txt")
	toWrite := fmt.Sprint(client1.File, ";",
		client1.NameAndLastName, ";",
		client1.DNI, ";",
		client1.CellPhone, ";",
		client1.Address, ";")
	file.WriteString(toWrite)

	toWrite = fmt.Sprint(client2.File, ";",
		client2.NameAndLastName, ";",
		client2.DNI, ";",
		client2.CellPhone, ";",
		client2.Address, ";")

	file.WriteString(toWrite)
	if err != nil {
		panic("Audodestrucción")
	}

	_, errR := os.ReadFile("cusmers.txt")
	defer func() {
		fmt.Println("Ejecución no finalizada")
	}()
	if errR != nil {
		panic("Archivo no encontrado o dañado")
	}
	recover()
	clientVerified, errVerified := validateClient(client1)
	defer func() {
		fmt.Println("Ejecución finalizada")
		fmt.Println("Se determinaron varios errores")
		fmt.Print("No quedaron archivos abiertos")
	}()
	if errVerified != nil {
		panic("Ingresaste valores nulos")
	}
	fmt.Println(clientVerified)
}
