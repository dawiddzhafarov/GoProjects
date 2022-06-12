package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	go task1(&wg)
	go task2(&wg)
	go task3(&wg)
	go task4(&wg)
	go task5(&wg)
	wg.Wait()
}

func task1(wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatalf("error %v", err)
	}
	fmt.Println("Task 1")
}

func task2(wg *sync.WaitGroup) {
	defer wg.Done()
	var count int
	for i := 0; i < 1000000; i++ {
		count += 1
	}
	fmt.Println("Task 2")
}
func task3(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task 3")
}
func task4(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 800)
	fmt.Println("Task 4")
}
func task5(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 500)
	fmt.Println("Task 5")
}
