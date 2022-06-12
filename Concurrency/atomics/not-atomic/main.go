package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type student struct {
	grades map[string]int
}

// altough it panics as there are multiple map writes, running it with -race flag
// doesnt show any race conditions
func main() {
	var wg sync.WaitGroup
	var val atomic.Value
	val.Store(student{grades: map[string]int{}})

	wg.Add(3)
	go func() {
		defer wg.Done()
		s := val.Load().(student)
		m := s.grades
		m["English"] = 10
		val.Store(student{grades: m})
		//s.grades["English"] = 10
	}()
	go func() {
		defer wg.Done()
		s := val.Load().(student)
		m := s.grades
		m["Math"] = 15
		val.Store(student{grades: m})
		//s.grades["Math"] = 15
	}()
	go func() {
		defer wg.Done()
		s := val.Load().(student)
		m := s.grades
		m["Physics"] = 2
		val.Store(student{grades: m})
		//s.grades["Physics"] = 2
	}()
	wg.Wait()

	s := val.Load().(student)
	fmt.Println(s)
}
