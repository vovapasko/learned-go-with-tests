package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	t.Run("test when everything goes well", func(t *testing.T) {
		data := "hello world"
		store := &SpyStore{response: data}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		got := response.Body.String()
		if got != data {
			t.Errorf("got %s, want %s", got, data)
		}

		if store.cancelled {
			t.Error("store should not be cancelled")
		}
	})
	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello world"
		store := &SpyStore{response: data}

		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		if !store.cancelled {
			t.Error("Store was not told to cancel")
		}

	})
}
