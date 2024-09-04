package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	channel3 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel1 <- "Channel 1 dan message"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		channel2 <- "Channel 2 dan message"
	}()
	go func() {
		time.Sleep(3 * time.Second)
		channel3 <- "Channel 3 dan message"
	}()

	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-channel1:
			fmt.Println("Received from", msg1)
		case msg2 := <-channel2:
			fmt.Println("Received from", msg2)
		case msg3 := <-channel3:
			fmt.Println("Received from", msg3)
		}
	}
}
