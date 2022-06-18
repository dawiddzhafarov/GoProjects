package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	t.Run("return faster respondse", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, but wanted %q", got, want)
		}

		if err != nil {
			t.Errorf("got an error but should't have")
		}
	})

	t.Run("return an error if any of the servers doesnt response withing 10s", func(t *testing.T) {
		serv := makeDelayedServer(50 * time.Millisecond)

		defer serv.Close()

		_, err := ConfigurableRacer(serv.URL, serv.URL, 20*time.Millisecond)

		if err == nil {
			t.Errorf("Expected to get and error, but didn't get one")
		}
	})

}
