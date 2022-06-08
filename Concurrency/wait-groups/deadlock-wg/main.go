package main

import "sync"

func main() {
	// wg waits for some goroutine that should call the Done() method,
	// but it never happens, so there is a deadlock-wg
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
