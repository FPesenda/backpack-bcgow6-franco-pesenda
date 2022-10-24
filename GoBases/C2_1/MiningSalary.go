package main

import "fmt"

const (
	A = 0.0
	B = 0.2
	C = 0.5
)

func salaryCalculator(salaryForHour, hours, extra float64) (salary float64) {
	salary = float64(salaryForHour * hours)
	salary += (salary * extra)
	return
}
func main() {
	salaryFH := 1200.0
	hours := 230.0
	extra := A
	fmt.Println("El salario al mes es: ", salaryCalculator(salaryFH, hours, extra))
}
