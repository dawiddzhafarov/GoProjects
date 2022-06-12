package main

import "sync/atomic"

// atomic value has to store the same type, otherwise causes panic

func main() {
	var v atomic.Value
	v.Store(1)
	v.Store("lol")
}
