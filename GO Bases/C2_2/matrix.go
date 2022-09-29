package main

import "fmt"

type matrix struct {
	Heigt     int
	Width     int
	Data      []float64
	Cuadratic bool
	MaxValue  float64
}

func (mat *matrix) Setter(numbers ...float64) {
	mat.Data = numbers
	if len(numbers) < (mat.Heigt * mat.Width) {
		for i := 0; i < ((mat.Heigt * mat.Width) - len(numbers)); i++ {
			mat.Data = append(mat.Data, 0.0)
		}
	}
}

func (mat matrix) Getter() {
	for i, v := range mat.Data {
		if i%mat.Width == 0 {
			fmt.Println()
		}
		fmt.Print(v, "\t ")
	}
	fmt.Println()
}

type Matrix interface {
	Setter(...float64)
	Getter()
}

func (m *matrix) NewMatrix() Matrix {
	return &matrix{}
}

func main() {

	m1 := matrix{
		Heigt:     3,
		Width:     3,
		Cuadratic: true,
		MaxValue:  10.0,
	}

	m1.Setter(3.0, 1.2, 4.5)
	m1.Getter()

}
