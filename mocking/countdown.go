package mocking

import (
	"fmt"
	"io"
	"time"
)

const countingAmounts = 3
const finalWord = "Go!"

func Counting(out io.Writer) {
	for i := countingAmounts; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(out, finalWord)
}
