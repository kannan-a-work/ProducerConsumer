package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	route := make(chan string)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go producer(route, i, &wg)
	}
	for i := 0; i < 10; i++ {
		go consumer(route, i)
	}
	wg.Wait()
	fmt.Println("All jobs completed")
}

func producer(route chan<- string, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	route <- fmt.Sprintf("Data from producer %d", id)
}
func consumer(route <-chan string, id int) {
	for d := range route {
		fmt.Println("Consumer", id, "received", d)
		time.Sleep(1 * time.Second) // Simulate processing time
	}
}
