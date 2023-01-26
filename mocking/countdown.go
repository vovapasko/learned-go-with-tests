package mocking

import (
	"fmt"
	"io"
)

const countingAmounts = 3
const finalWord = "Go!"

func Counting(out io.Writer, sleeper Sleeper) {
	for i := countingAmounts; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}
