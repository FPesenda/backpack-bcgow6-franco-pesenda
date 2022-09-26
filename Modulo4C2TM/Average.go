package main

import (
	"errors"
	"fmt"
)

func average(rating ...float64) (averageStudentRating float64, err error) {
	sumRating := 0.0
	for _, v := range rating {
		if v < 0 {
			err = errors.New("No puede ingresar un valor negativo")
			averageStudentRating = 0
			return
		}
		sumRating += v
	}
	averageStudentRating = sumRating / float64(len(rating))
	return

}

func main() {
	averageStudent, err := average(4, 2, -3, 3, 4, 5, 6, 10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El promedio de los alumnos es: ", averageStudent)
	}

}
