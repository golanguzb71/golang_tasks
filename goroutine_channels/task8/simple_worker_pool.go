package main

import (
	"fmt"
	"sync"
)

/*
Oddiy Task
Ma'lumotlarni qayta ishlash

Vazifa: Worker pool bilan kichik ma'lumotlar to‘plamini (masalan, integerlar ro‘yxati) qayta ishlang. Har bir worker elementlarni qayta ishlaydi, masalan, ularning kvadratini hisoblash.
Maqsad: Worker pool bilan parallel ma'lumotlarni qayta ishlashni tushunish.
*/

func main() {
	jobs := make(chan int, 6)
	results := make(chan int, 6)
	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for i := 1; i <= 6; i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println("Result :", result)
	}
}

func worker(w int, jobs chan int, results chan int, s *sync.WaitGroup) {
	defer s.Done()
	for job := range jobs {
		fmt.Printf("Worker %d proccessing %d \n", w, job)
		results <- job * job
	}
}
