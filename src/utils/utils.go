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
