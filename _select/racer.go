package _select

import "net/http"

func Racer(firstUrl, secondUrl string) (winner string) {
	select {
	case <-ping(firstUrl):
		return firstUrl
	case <-ping(secondUrl):
		return secondUrl
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
