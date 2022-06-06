package main

import (
	"fmt"
	"time"
)

type object struct {
	name   string
	number int
}

func main() {

	c := make(chan object)
	go count2("exe", c)

	for msg := range c {
		fmt.Println(msg)
	}

}

func count2(thing string, c chan object) {
	for i := 0; i < 5; i++ {
		c <- object{thing, i}
		time.Sleep(time.Millisecond * 1000)
	}
	close(c)
}
