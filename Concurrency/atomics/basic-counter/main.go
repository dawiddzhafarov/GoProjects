package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var count int64
	var wg sync.WaitGroup
	wg.Add(1)

	//reader
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 10)
		fmt.Println("count in goroutine", atomic.LoadInt64(&count))
	}()

	wg.Add(50)
	//writers
	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * 10)
			atomic.AddInt64(&count, 1)
		}()
	}
	wg.Wait()
	fmt.Println("count in main", count)
}
