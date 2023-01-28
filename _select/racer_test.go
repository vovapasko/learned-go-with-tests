package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("test both valid servers", func(t *testing.T) {
		slowServer := makeServer(20 * time.Millisecond)
		fastServer := makeServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		got, _ := Racer(slowURL, fastURL)
		want := fastURL

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("test both servers with > 10 seconds timeout", func(t *testing.T) {
		server1 := makeServer(11 * time.Millisecond)
		server2 := makeServer(12 * time.Millisecond)

		defer server1.Close()
		defer server2.Close()

		url1 := server1.URL
		url2 := server2.URL

		_, err := Racer(url1, url2)
		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})

}

func makeServer(delayTime time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delayTime)
		w.WriteHeader(http.StatusOK)
	}))
}
