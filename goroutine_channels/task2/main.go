package main

import (
	"fmt"
	"sync"
)

/*
Kanallar bilan ishlash: Go-routine orqali kanalda sonlar yuboring va asosiy funksiya bu sonlarni o'qib ekranga chiqarsin.

Maqsad: Kanallar orqali ma'lumot almashish mexanizmini tushunish.
*/

func main() {
	numbersChan := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go sendNumbersToMain(&wg, numbersChan)
	go func() {
		for number := range numbersChan {
			fmt.Println("received : ", number)
		}
	}()
	wg.Wait()
	close(numbersChan)
}

func sendNumbersToMain(wg *sync.WaitGroup, numbersChan chan int) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		numbersChan <- i
	}
}
