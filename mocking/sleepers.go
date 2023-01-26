package mocking

import "time"

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	CallsAmount int
}

func (sleeper *SpySleeper) Sleep() {
	sleeper.CallsAmount += 1
}

type DefaultSleeper struct{}

func (sleeper *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type CountdownOperationsSpy struct {
	Operations []string
}

func (operationsList *CountdownOperationsSpy) Sleep() {
	operationsList.Operations = append(operationsList.Operations, sleepOperationSignature)
}

func (operationsList *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	operationsList.Operations = append(operationsList.Operations, writeOperationSignature)
	return
}

const writeOperationSignature = "write"
const sleepOperationSignature = "sleep"
