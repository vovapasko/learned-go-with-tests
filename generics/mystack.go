package generics

import "errors"

type MyStack[T any] struct {
	values []T
}

func (stack *MyStack[T]) IsEmpty() bool {
	return len(stack.values) == 0
}

func (stack *MyStack[T]) Push(value T) {
	stack.values = append(stack.values, value)
}

func (stack *MyStack[T]) Pop() (T, error) {
	if stack.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}
	lastIndex := len(stack.values) - 1
	stackTopValue := stack.values[lastIndex]
	stack.values = stack.values[:lastIndex]
	return stackTopValue, nil
}
