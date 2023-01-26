package mocking

import (
	"fmt"
	"io"
	"time"
)

const countingAmounts = 3
const finalWord = "Go!"

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	CallsAmount int
}

type DefaultSleeper struct{}

func (sleeper *SpySleeper) Sleep() {
	sleeper.CallsAmount += 1
}

func (sleeper *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Counting(out io.Writer, sleeper Sleeper) {
	for i := countingAmounts; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}
