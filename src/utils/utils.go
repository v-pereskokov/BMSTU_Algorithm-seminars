package utils

import "errors"

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

func Reverse(numbers []int, begin int, end int) ([]int, error) {
	if begin == -1 {
		begin = 0
	}

	if end != -1 {
		end = len(numbers) - 1
	}

	if begin > end {
		return numbers, errors.New("Begin more than end")
	}

	for i, j := begin, end; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	return numbers, nil
}
