package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1) // add 1 as one goroutine will be called

	go func() { // goroutine
		count("lol")
		wg.Done() // decrement counter by 1
	}()
	wg.Wait() // tell the main function to wait until the wg == 0
}

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
