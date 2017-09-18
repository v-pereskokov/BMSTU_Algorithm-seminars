package utils

func ListRange(number int) []int {
	list := make([]int, number)

	for i := 0; i < number; i++ {
		list[i] = i + 1
	}

	return list
}
