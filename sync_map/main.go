package main

import (
	"fmt"
	"sync"
	"time"
)

// Order represents a food order
type Order struct {
	ID       string
	Customer string
	Items    []string
}

// Database simulation
func processOrder(order Order) {
	fmt.Printf("Processing order ID: %s for customer: %s\n", order.ID, order.Customer)
	time.Sleep(1 * time.Second) // Simulate DB update delay
	fmt.Printf("Order ID: %s processed successfully!\n", order.ID)
}

func main() {
	orderQueue := make(chan Order, 100)
	var orderLocks sync.Map
	var wg sync.WaitGroup

	// Worker function to process orders
	worker := func() {
		for order := range orderQueue {
			mutexInterface, _ := orderLocks.LoadOrStore(order.ID, &sync.Mutex{})
			mutex := mutexInterface.(*sync.Mutex)

			mutex.Lock() // Ensure only one goroutine updates the DB per order
			processOrder(order)
			mutex.Unlock()
		}
		wg.Done()
	}

	// Start worker pool
	numWorkers := 5
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker()
	}

	// Simulate incoming orders
	orders := []Order{
		{"1", "Alice", []string{"Burger", "Fries"}},
		{"2", "Bob", []string{"Pizza"}},
		{"3", "Charlie", []string{"Pasta", "Salad"}},
		{"1", "Alice", []string{"Burger", "Fries"}}, // Duplicate Order ID (should not be processed twice simultaneously)
		{"4", "David", []string{"Sushi"}},
	}

	for _, order := range orders {
		orderQueue <- order
	}

	close(orderQueue)
	wg.Wait() // Wait for all workers to finish
}
