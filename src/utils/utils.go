package utils

import "errors"

func Range(begin int, end int) ([]int, error) {
	if begin > end {
		return make([]int, 0), errors.New("Begin more than end")
	}

	if end < 1 {
		return make([]int, 0), errors.New("End less than 1")
	}

	length := end - begin
	list := make([]int, length)

	for i := 0; i < length; i++ {
		list[i] = begin + i
	}

	return list, nil
}

func Reverse(numbers []int, begin int, end int) ([]int, error) {
	if begin > end {
		return numbers, errors.New("Begin more than end")
	}

	if begin == -1 {
		begin = 0
	}

	if end == -1 {
		end = len(numbers) - 1
	}

	for i, j := begin, end; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	return numbers, nil
}
