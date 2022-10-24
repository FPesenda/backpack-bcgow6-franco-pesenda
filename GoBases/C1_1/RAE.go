package main

import (
	"fmt"
)

func main() {

	palabra := "tomate"
	sum := 0
	var letras []string
	for i := 0; i < len(palabra); i++ {
		letras = append(letras, string(palabra[i]))
		sum++
	}
	fmt.Print("La cantidad total de letras es: ", sum, "\n")
	for i := 0; i < len(letras); i++ {
		fmt.Println(letras[i])
	}

}
