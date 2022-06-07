package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		custom()
		wg.Done()
	}()

	wg.Wait()

}
func custom() {
	var wg sync.WaitGroup
	wg.Add(1)
	for counter := 10; counter >= 0; counter-- {
		fmt.Println(counter)
		time.Sleep(time.Millisecond * 1000)
	}
	go func() {
		inner()
		wg.Done()
	}()
	wg.Wait()
}

func inner() {
	for j := 0; j <= 5; j++ {
		fmt.Println(j)
		time.Sleep(time.Millisecond * 500)
	}
}
