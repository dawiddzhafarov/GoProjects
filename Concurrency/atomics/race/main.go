package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int32
	var wg sync.WaitGroup
	wg.Add(5)
	go func() {
		defer wg.Done()
		count += 10
	}()
	go func() {
		defer wg.Done()
		count -= 15
	}()
	go func() {
		defer wg.Done()
		count += 1
	}()
	go func() {
		defer wg.Done()
		count = 0
	}()
	go func() {
		defer wg.Done()
		count = 1011
	}()
	wg.Wait()
	fmt.Print("count", count)
}
