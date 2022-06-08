package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go work(&wg, i+1)
	}
	wg.Wait()                                     // wait until 10 gorouitnes are finished
	fmt.Println("elapsed time:", time.Since(now)) // around 100ms
	fmt.Println("main completed")

}

func work(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 100)
	fmt.Println("task", id, " is done")
}
