package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	empleadoNombre := "Benjamin"
	fmt.Print(employees[empleadoNombre], "\n")
	//CUANTOS TIENEN MAS DE 21 AÑOS
	sum := 0
	for _, v := range employees {
		if v > 21 {
			sum++
		}
	}
	fmt.Print("La cantidad de empleados mayores a 21 son: ", sum, "\n")
	employees["Federico"] = 25
	delete(employees, "Pedro")
	for k, v := range employees {
		fmt.Println(k, " empleado con ", v, " años.")
	}
}
