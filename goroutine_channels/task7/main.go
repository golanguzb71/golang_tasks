package main

import (
	"fmt"
	"sync"
)

func main() {
	resultChan := make(chan string, 4)
	var wg sync.WaitGroup
	a := 12.0
	b := 4.0

	wg.Add(4)
	go Add(a, b, resultChan, &wg)
	go Subtract(a, b, resultChan, &wg)
	go Multiply(a, b, resultChan, &wg)
	go Divide(a, b, resultChan, &wg)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		fmt.Println(result)
	}
}

func Divide(a float64, b float64, resultChan chan<- string, s *sync.WaitGroup) {
	defer s.Done()
	result := a + b
	resultChan <- fmt.Sprintf("%f + %f = %f", a, b, result)
}

func Multiply(a float64, b float64, resultChan chan string, s *sync.WaitGroup) {
	defer s.Done()
	result := a - b
	resultChan <- fmt.Sprintf("%f * %f = %f", a, b, result)
}

func Subtract(a float64, b float64, resultChan chan string, s *sync.WaitGroup) {
	defer s.Done()
	result := a - b
	resultChan <- fmt.Sprintf("%f -%f=%f", a, b, result)
}

func Add(a float64, b float64, resultChan chan string, s *sync.WaitGroup) {
	defer s.Done()
	if b != 0 {
		result := a / b
		resultChan <- fmt.Sprintf("%f / %f = %f", a, b, result)
	} else {
		resultChan <- "Error: Division by zero"
	}
}
