package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciHappy(t *testing.T) {
	//arange
	size := 10
	expectedResult := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	//act
	numbers := fibonacci(size)
	//asert
	assert.Equal(t, expectedResult, numbers)
}
