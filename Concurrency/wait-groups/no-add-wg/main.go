package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// there is no Add(), so wg doesn't wait for the execution of a goroutine
	var wg sync.WaitGroup
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 10000)
		fmt.Println("goroutine done")
	}()
	wg.Wait()
	fmt.Println("Executed immediately")
}
