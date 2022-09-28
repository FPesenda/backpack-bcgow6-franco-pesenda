package main

import "fmt"

func main() {
	salary := 1222222000
	minimunSalary := 150000
	if salary < minimunSalary {
		errString := fmt.Errorf("Error: el minimo imponible es %d y el salario ingresado es %d", minimunSalary, salary)
		fmt.Println(errString)
	} else {
		fmt.Println("Debe pagar impuestos")
	}
}
