package main

import (
	"fmt"
)

type TaxError struct {
	Message   string
	CodeError int
}

func (err *TaxError) Error() string {
	return fmt.Sprint("error: el salario ingresado no alcanza el mínimo imponible")
}

func validateMinimungSalary(salary int) (message string, err error) {
	if salary < 150000 {
		err = &TaxError{Message: message, CodeError: 500}
	} else {
		message = "Debe pagar impuestos"
	}
	return
}

func main() {
	salary, err := validateMinimungSalary(151000)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(salary)
	}
}
