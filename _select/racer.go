package _select

import "net/http"

func Racer(firstUrl, secondUrl string) (winner string, err error) {
	select {
	case <-ping(firstUrl):
		return firstUrl, nil
	case <-ping(secondUrl):
		return secondUrl, nil
	}
}

func ping(url string) chan struct{} {
	chanel := make(chan struct{})
	go func() {
		http.Get(url)
		close(chanel)
	}()

	return chanel
}
