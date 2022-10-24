package main

import (
	"errors"
	"fmt"
)

func minimun(rating ...float64) (min float64) {
	for i, v := range rating {
		if i == 0 {
			min = v
		}
		if v < min {
			min = v
		}
	}
	return min
}

func averages(rating ...float64) (avg float64) {
	ratingTotal := 0.0
	for _, v := range rating {
		ratingTotal += v
	}
	avg = ratingTotal / float64(len(rating))
	return
}

func maximun(rating ...float64) (max float64) {
	for i, v := range rating {
		if i == 0 {
			max = v
		}
		if v > max {
			max = v
		}
	}
	return max
}

func selectOperation(operation string) (functionSelected func(...float64) float64, err error) {
	switch operation {
	case "minimun":
		functionSelected = minimun
		return
	case "maximun":
		functionSelected = maximun
		return
	case "average":
		functionSelected = averages
		return
	}
	err = errors.New("No ingreso un operador soportado")
	return
}

func main() {
	operation, err := selectOperation("averasdasdage")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(operation(3, 4, 2, 3, 7, 8, 9))
	}
}
