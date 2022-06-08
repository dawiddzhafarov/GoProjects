package main

import (
	"sync"
	"time"
)

func main() {
	// reusing wg that hasn't returned its Wait() method causes panic
	// should return the Wait(), and then we can reuse it
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond * 1000)
		wg.Done()
		wg.Add(1)
	}()

	wg.Wait()
}
