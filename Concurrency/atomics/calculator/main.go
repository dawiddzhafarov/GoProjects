package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Calculator struct {
	res atomic.Value
}

func newCalculator() Calculator {
	c := Calculator{}
	c.res.Store(float64(0))
	return c
}
func (c *Calculator) add(n float64) {
	c.res.Store(c.result() + n)
}
func (c *Calculator) sub(n float64) {
	c.res.Store(c.result() - n)
}
func (c *Calculator) mul(n float64) {
	c.res.Store(c.result() * n)
}
func (c *Calculator) div(n float64) {
	if n == 0 {
		panic("Cant divide by 0")
	}
	c.res.Store(c.result() / n)
}
func (c *Calculator) result() float64 {
	r, ok := c.res.Load().(float64)
	if !ok {
		panic("Wrong type")
	}
	return r
}

func main() {
	var wg sync.WaitGroup
	c := newCalculator()
	wg.Add(5)
	go func() {
		defer wg.Done()
		c.add(10)
	}()
	go func() {
		defer wg.Done()
		c.sub(23)
	}()
	go func() {
		defer wg.Done()
		c.mul(120)
	}()
	go func() {
		defer wg.Done()
		c.mul(1)
	}()
	go func() {
		defer wg.Done()
		c.div(34)
	}()
	wg.Wait()
	fmt.Println("result:", c.res)
}
