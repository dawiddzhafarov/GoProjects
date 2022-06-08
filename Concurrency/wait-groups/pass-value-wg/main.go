package main

import (
	"fmt"
	"sync"
)

func main() {
	// wg should be passed as a pointer, not a value
	var wg sync.WaitGroup
	wg.Add(1)
	go work(wg)
	wg.Wait()
}

func work(wg sync.WaitGroup) { // indication that wg is a value, not a pointer
	defer wg.Done()
	fmt.Println("work done")
}
