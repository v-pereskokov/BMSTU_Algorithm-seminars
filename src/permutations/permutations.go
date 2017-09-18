package permutations

import (
	"github.com/BMSTU_Algorithm-seminars/src/utils"
)

func Permutations(number int) <-chan []int {
	c := make(chan []int)

	go func() {
		defer close(c)
		list, _ := utils.Range(1, number+1)

		for {
			c <- list

			i := number - 2
			for i >= 0 && list[i] >= list[i+1] {
				i--
			}

			if i < 0 {
				return
			}

			min := i + 1

			for j := i + 2; j < number; j++ {
				if list[j] > list[i] && list[j] < list[min] {
					min = j
				}
			}

			list[i], list[min] = list[min], list[i]

			utils.Reverse(list, i+1, number-1)
		}
	}()

	return c
}
