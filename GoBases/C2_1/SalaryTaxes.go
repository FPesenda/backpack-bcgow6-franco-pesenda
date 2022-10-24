package main

import "fmt"

func taxCalculate(salary float64) (salaryWithTaxes float64) {
	taxes := 0.0
	if salary > 50000 {
		taxes += 0.17
	}
	if salary > 150000 {
		taxes += 0.1
	}
	salaryWithTaxes = salary * (1 - taxes)
	return
}

func main() {
	salary := 40000.0
	fmt.Println("El salario ", salary, " termino siendo: ", taxCalculate(salary))
}
