package main

import "fmt"

func main() {
	edad := 30
	empleado := true
	antiguedad := 1
	sueldo := 90000
	entregar := true

	if edad < 22 {
		entregar = false
	}
	if !empleado {
		entregar = false
	}
	if antiguedad <= 1 {
		entregar = false
	}

	if entregar {
		if sueldo > 100000 {
			fmt.Println("Prestamo otorgado sin intereses")
		} else {
			fmt.Println("Prestamo otorgado con intereses")
		}
	} else {
		fmt.Println("Prestamo no otorgado")
	}

}
