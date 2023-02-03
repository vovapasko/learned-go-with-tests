package arrays_slices_with_generics

type Transaction struct {
	From, To string
	Sum      int
}

func Reduce[T any](collection []T, function func(a, b T) T, initialValue T) T {
	result := initialValue
	for _, number := range collection {
		result = function(result, number)
	}
	return result
}

func BalanceFor(transactions []Transaction, name string) int {
	personBalance := 0
	for _, transaction := range transactions {
		if transaction.From == name {
			personBalance -= transaction.Sum
		}
		if transaction.To == name {
			personBalance += transaction.Sum
		}
	}
	return personBalance
}
