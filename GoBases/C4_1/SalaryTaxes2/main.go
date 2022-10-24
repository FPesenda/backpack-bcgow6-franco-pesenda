package main

import (
	"errors"
	"fmt"
)

func validateMinimungSalary(salary int) (message string, err error) {
	if salary < 150000 {
		err = errors.New("error: el salaro ingresado no alcanza el minimo imponible")
	} else {
		message = "Debe pagar impuestos"
	}
	return
}

func main() {
	salary, err := validateMinimungSalary(10000000)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(salary)
	}
}
