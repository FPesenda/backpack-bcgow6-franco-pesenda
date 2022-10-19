package gotesting

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHappySubstractPositiveResult(t *testing.T) {
	number1 := 10
	number2 := 5
	numberExpected := 5
	result := Substract(number1, number2)
	assert.Equal(t, result, numberExpected, "deben ser iguales")
}

func TestHappySubstractNegativeResult(t *testing.T) {
	number1 := 5
	number2 := 10
	numberExpected := -5
	result := Substract(number1, number2)
	assert.Equal(t, result, numberExpected, "deben ser iguales")
}

func TestHappyOrderNumbers(t *testing.T) {
	numbers := []int{3, 4, 5, 1}
	orderExpected := []int{1, 3, 4, 5}
	result := OrderNumbersDesc(numbers)
	assert.Equal(t, result, orderExpected, "deben ser guales")
}

func TestHappyDivide(t *testing.T) {
	number1 := 10
	number2 := 2
	numberExpected := 5
	result, err := Divide(number1, number2)
	assert.Nil(t, err, "debe el error ser nulo")
	assert.Equal(t, result, numberExpected, "deben ser iguales")
}

func TestSadDivide(t *testing.T) {
	number1 := 10
	number2 := 0
	errExpected := errors.New("el divisor no puede ser 0")
	_, err := Divide(number1, number2)
	assert.NotNil(t, err, "no debe ser nulo")
	assert.Equal(t, err, errExpected, "deben ser iguales")
}
