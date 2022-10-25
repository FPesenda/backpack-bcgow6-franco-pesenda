package fibonacci

func main() {

}

func fibonacci(size int) (result []int) {
	num1 := 0
	num2 := 1
	result = []int{0, 1}
	for i := 0; i < (size - 2); i++ {
		num2 = num1 + num2
		num1 = num2 - num1
		result = append(result, num2)
	}
	return
}
