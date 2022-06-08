package main

import "sync"

func main() {
	// for is so fast that the next iteration tries to reuse the wg by adding
	// causing panic, as not all goroutines are finished and Wait() doesn't return
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		go func() {
			wg.Add(3)
			go func() {
				wg.Done()
			}()
			go func() {
				wg.Done()
			}()
			go func() {
				wg.Done()
			}()
			wg.Wait()
		}()
	}
}
