package main

import (
	"fmt"
	"time"
)

func main() {
	// without wg or channels, main may finish before some
	// of the goroutines are finished
	for i := 0; i < 10; i++ {
		go work(i + 1)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("main completed")
}

func work(id int) {
	time.Sleep(time.Millisecond * 100)
	fmt.Println("task", id, " is done")
}
