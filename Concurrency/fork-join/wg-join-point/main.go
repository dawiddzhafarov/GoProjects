package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	var wg sync.WaitGroup
	wg.Add(1) // add 1 goroutine

	go func() {
		defer wg.Done()
		work() // fork point
	}()

	wg.Wait() // join point; wait until the counter == 0
	fmt.Println("elapsed: ", time.Since(now))
	fmt.Print("done waiting, main exits\n")
}

func work() {
	time.Sleep(time.Millisecond * 500)
	fmt.Println("work done")
}
