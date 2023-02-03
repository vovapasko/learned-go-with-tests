package arrays_slices_with_generics

func Reduce[T any](collection []T, function func(a, b T) T, initialValue T) T {
	result := initialValue
	for _, number := range collection {
		result = function(result, number)
	}
	return result
}
