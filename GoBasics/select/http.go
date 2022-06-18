package _select

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecTimeout = 10 * time.Second

func Racer(url1, url2 string) (winner string, err error) {
	return ConfigurableRacer(url1, url2, tenSecTimeout)
}

func ConfigurableRacer(url1, url2 string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for responses of %s and %s", url1, url2)
	}
}

func ping(s string) chan struct{} {
	ch := make(chan struct{}) // struct does not allocate any memory, compared to bool for example
	// use make to create channels, as var ch chan struct{] results in creating nil channel
	// and sending to a nil channels causes it to block
	go func() {
		http.Get(s)
		close(ch) // closing channel == sending information, sending info that the channel is closed
	}()
	return ch
}

func measureResponseTime(s string) time.Duration {
	start := time.Now()
	http.Get(s)
	return time.Since(start)
}
