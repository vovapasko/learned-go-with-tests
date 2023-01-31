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
	t         *testing.T
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func assertCancelled(s *SpyStore) {
	s.t.Helper()
	if !s.cancelled {
		s.t.Error("store should not be cancelled")
	}

}

func assertNotCancelled(s *SpyStore) {
	s.t.Helper()
	if s.cancelled {
		s.t.Error("Store should be canceled")
	}
}

func TestServer(t *testing.T) {
	t.Run("test when everything goes well", func(t *testing.T) {
		data := "hello world"
		store := &SpyStore{response: data, t: t}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		got := response.Body.String()
		if got != data {
			t.Errorf("got %s, want %s", got, data)
		}

		assertNotCancelled(store)
	})
	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello world"
		store := &SpyStore{response: data, t: t}

		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertCancelled(store)
	})
}
