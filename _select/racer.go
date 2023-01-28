package _select

import (
	"fmt"
	"net/http"
	"time"
)

const defaultTimeout = 10 * time.Second

func ConfigurableRacer(firstUrl, secondUrl string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(firstUrl):
		return firstUrl, nil
	case <-ping(secondUrl):
		return secondUrl, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("time is out")
	}
}

func Racer(firstUrl, secondUrl string) (winner string, err error) {
	return ConfigurableRacer(firstUrl, secondUrl, defaultTimeout)
}

func ping(url string) chan struct{} {
	chanel := make(chan struct{})
	go func() {
		http.Get(url)
		close(chanel)
	}()

	return chanel
}
