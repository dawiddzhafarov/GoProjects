package main

import "sync"

func main() {
	// negative counter, causes panic
	var wg sync.WaitGroup
	wg.Done()
}
