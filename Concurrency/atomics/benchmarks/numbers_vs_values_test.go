package benchmarks

import (
	"sync"
	"sync/atomic"
	"testing"
)

// using functions for types is slightly faster than using atomic.Value
func BenchmarkStoreInt64(b *testing.B) {
	var wg sync.WaitGroup
	wg.Add(b.N)
	var count int64
	for i := 0; i < b.N; i++ {
		go func(i int) {
			defer wg.Done()
			atomic.StoreInt64(&count, int64(i))
		}(i)
	}
	wg.Wait()
}

func BenchmarkStoreValue(b *testing.B) {
	var wg sync.WaitGroup
	var v atomic.Value
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func(i int) {
			defer wg.Done()
			v.Store(int64(i))
		}(i)
	}
	wg.Wait()
}
