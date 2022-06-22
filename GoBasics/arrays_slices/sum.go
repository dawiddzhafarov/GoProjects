package main

func Sum(values []int) int {
	/*
		sum := 0
		for _, num := range values {
			sum += num
		}
		return sum

	*/
	add := func(acc, x int) int {
		return acc + x
	}
	return Reduce(values, add, 0)
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(slices ...[]int) []int {
	/*
		var sums []int

		for _, numbers := range slices {
			if len(numbers) == 0 {
				sums = append(sums, 0)
			} else {
				sums = append(sums, Sum(numbers[1:]))
			}

		}
	*/
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}
	return Reduce(slices, sumTail, []int{})
}

func Reduce[A any](collection []A, accumulator func(A, A) A, initialValue A) A {
	var result = initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}
	return result
}
