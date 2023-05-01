package arrays

func SumAll(numbersToSum ...[]int) []int {

	var sum []int

	for _, numbers := range numbersToSum {
		sum = append(sum, ArraySum(numbers))
	}

	return sum
}

func SumAllTrails(numbersToSum ...[]int) []int {

	var sum []int

	for _, numbers := range numbersToSum {

		if len(numbers) == 0 {
			sum = append(sum, 0)
		} else {
			tail := numbers[1:]
			sum = append(sum, ArraySum(tail))
		}
	}

	return sum
}
