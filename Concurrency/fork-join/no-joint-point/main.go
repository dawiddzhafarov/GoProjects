package main

import (
	"fmt"
	"time"
)

func main() {
	go work() // fork point
	// but main doesn't wait  for this goroutine to finish, it exits
	// printing only the last message
	time.Sleep(time.Millisecond * 100)
	fmt.Print("done waiting, main exits")
}

func work() {
	time.Sleep(time.Millisecond * 500)
	fmt.Println("work done")
}
