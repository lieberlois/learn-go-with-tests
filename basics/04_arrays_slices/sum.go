package arraysslices

// Sum returns the sum of all integers in an array
func Sum(numbers []int) (sum int) {
	for _, elem := range numbers {
		sum += elem
	}
	return
}

// SumAll returns an array of sums 
func SumAll(arraysToSum ...[]int) (sums []int) {
	for _, elem := range arraysToSum {
		sums = append(sums, Sum(elem))
	}
	return
}

// SumAllTails returns each sum of all tails 
func SumAllTails(arraysToSum ...[]int) (sums []int) {
	for _, elem := range arraysToSum {
		if len(elem) == 0{
			sums = append(sums, 0)
		} else {
			tail := elem[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}