package add_benchmark

import (
	"sync"
	"testing"
)

func BenchmarkAddOne(b *testing.B) {
	var count int
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1) // add one each iteration
		go func() {
			defer wg.Done()
			count++
		}()
	}
	wg.Wait()
}

func BenchmarkAddMany(b *testing.B) {
	var count int
	var wg sync.WaitGroup
	wg.Add(b.N) // add number of tasks to do in one go, depends on testing
	// framework (what number it is going to be)
	for i := 0; i < b.N; i++ {
		go func() {
			defer wg.Done()
			count++
		}()
	}
	wg.Wait()
}
