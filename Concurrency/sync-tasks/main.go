package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// main runs in its own goroutine, exits straight away
	// we have to tell main func to wait until others tasks are finished
	// remove go keyword, longer time to execute
	done := make(chan struct{})
	go task1(done)
	go task2(done)
	go task3(done)
	go task4(done)
	<-done
	<-done
	<-done
	<-done
	fmt.Println("Elapsed:", time.Since(now))
}

func task1(done chan struct{}) {
	time.Sleep(time.Millisecond * 200)
	fmt.Println("Task 1 completed")
	done <- struct{}{}
}

func task2(done chan struct{}) {
	time.Sleep(time.Millisecond * 5000)
	fmt.Println("Task 2 completed")
	done <- struct{}{}
}

func task3(done chan struct{}) {
	time.Sleep(time.Millisecond)
	fmt.Println("Task 3 completed")
	done <- struct{}{}
}

func task4(done chan struct{}) {
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task 4 completed")
	done <- struct{}{}
}
