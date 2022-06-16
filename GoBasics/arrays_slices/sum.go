package main

func Sum(values []int) int {
	sum := 0
	for _, num := range values {
		sum += num
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(slices ...[]int) []int {
	var sums []int

	for _, numbers := range slices {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers[1:]))
		}

	}
	return sums
}
