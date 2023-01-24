package arrays_slices

func Sum(numbers []int) (sum int) {
	sum = 0
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	var summedArray []int
	for _, collection := range numbersToSum {
		summedArray = append(summedArray, Sum(collection))
	}
	return summedArray
}
