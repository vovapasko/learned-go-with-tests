package arrays_slices

func Sum(numbers [5]int) (sum int) {
	sum = 0
	for _, number := range numbers {
		sum += number
	}
	return
}
