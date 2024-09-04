package main

import (
	"fmt"
	"sync"
)

/*
Fan-Out Fan-In Pattern: 10 ta Go-routinedan iborat bo'lgan ishchilarni (workers) yarating, ular kanal orqali topshiriq qabul qilib, natijalarni bitta umumiy kanalda to'plab, natijalarni ekranga chiqaradigan boshqaruvchini (coordinator) yarating.

Maqsad: Go-routinelar va kanallar bilan bir nechta parallel jarayonlarni boshqarish
*/

const numWorkers = 10

func main() {
	tasks := make(chan int, 100)
	results := make(chan int, 100)

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	go func() {
		for i := 1; i <= 100; i++ {
			tasks <- i
		}
		close(tasks)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("Result:", result)
	}
}

func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		results <- task * 2
		fmt.Printf("Worker %d processed task %d\n", id, task)
	}
}
