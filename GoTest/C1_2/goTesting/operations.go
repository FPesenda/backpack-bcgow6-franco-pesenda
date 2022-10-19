package gotesting

import (
	"errors"
	"sort"
)

func Substract(nunmber1, nunmber2 int) (result int) {
	result = nunmber1 - nunmber2
	return
}

func Divide(number1, number2 int) (result int, err error) {
	if number2 == 0 {
		err = errors.New("el divisor no puede ser 0")
		return
	}
	result = number1 / number2
	return
}

func OrderNumbersDesc(numbers []int) (result []int) {
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	result = numbers
	return
}
