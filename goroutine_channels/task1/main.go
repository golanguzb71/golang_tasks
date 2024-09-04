package main

import (
	"fmt"
	"strconv"
	"sync"
)

/*
Oddiy Go-routine: Bitta Go-routinedan foydalangan holda ekranga "Hello, World!" so'zini 10 marta chop eting.

Maqsad: Go-routine yaratish va uni qanday ishga tushirishni tushunish.
*/

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go printHelloWorld(&wg)
	wg.Wait()
}

func printHelloWorld(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println(strconv.Itoa(i) + " ) Hello world")
	}
}
