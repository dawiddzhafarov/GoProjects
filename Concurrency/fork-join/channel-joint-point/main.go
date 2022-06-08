package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	done := make(chan struct{})

	go func() {
		work()
		done <- struct{}{}
	}()

	<-done

	fmt.Println("elapsed: ", time.Since(now))
	fmt.Print("done waiting, main exits\n")
}

func work() {
	time.Sleep(time.Millisecond * 500)
	fmt.Println("work done")
}
