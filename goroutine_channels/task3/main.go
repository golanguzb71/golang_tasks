package main

import (
	"fmt"
	"sync"
)

/*
Buffered Channel: Buffered channel yaratib, unga 5 ta son yuboring, keyin ularni qabul qiling va ekranga chop eting.

Maqsad: Buffered channel qanday ishlashini tushunish.
*/

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	bufferChannel := make(chan int, 5)
	go numberSender(bufferChannel, &wg)
	go func() {
		defer wg.Done()
		for number := range bufferChannel {
			fmt.Println("Received : ", number)
		}
	}()
	wg.Wait()
}

func numberSender(bufferChannel chan int, s *sync.WaitGroup) {
	defer s.Done()
	for i := 0; i < 10; i++ {
		bufferChannel <- i
	}
	close(bufferChannel)
}
