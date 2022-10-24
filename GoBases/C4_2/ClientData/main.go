package main

import (
	"fmt"
	"os"
)

func readCusmoers() (fil []byte, err error) {

	fil, err = os.ReadFile("ctomers.txt")
	return
}

func main() {
	filReaded, errReaded := readCusmoers()
	defer func() {
		fmt.Println("ejecución finalizada")
	}()
	if errReaded != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	fmt.Println(string(filReaded))

}
