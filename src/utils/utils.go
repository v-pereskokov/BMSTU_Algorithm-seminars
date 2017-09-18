package utils

func Range(number int) []int {
	if number < 1 {
		return make([]int, 0)
	}

	list := make([]int, number)

	for i := 0; i < number; i++ {
		list[i] = i + 1
	}

	return list
}

func Reverse(numbers []int) []int {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	return numbers
}
