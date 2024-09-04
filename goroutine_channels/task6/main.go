package main

import (
	"fmt"
	"time"
)

/*
Go-routine ishlashini kutib, lekin ma'lum bir vaqtdan keyin kutishni to'xtatib, vaqt tugash mexanizmini amalga oshiring.

Maqsad: Go-routinelar va vaqtni sinchkovlik bilan boshqarish.
*/

func main() {
	resultChannel := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		resultChannel <- "Go routine completed"
	}()
	select {
	case result := <-resultChannel:
		fmt.Println(result)
	case <-time.After(1*time.Second + 500*time.Millisecond):
		fmt.Println("Timeout! Go-routine took too long.")
	}
}
