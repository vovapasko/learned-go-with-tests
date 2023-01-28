package _select

import (
	"net/http"
	"time"
)

func Racer(firstUrl, secondUrl string) (winner string) {
	timeFirstRequest := measureRequestTime(firstUrl)
	timeSecondRequest := measureRequestTime(secondUrl)

	if timeFirstRequest > timeSecondRequest {
		return secondUrl
	}

	return firstUrl
}

func measureRequestTime(url string) time.Duration {
	startFirstRequest := time.Now()
	_, _ = http.Get(url)
	return time.Since(startFirstRequest)
}
