package main

import (
	"fmt"
	"time"
)

func printMessage(message string, channel chan string) {
	for i := 0; i < 10; i++ {
		fmt.Println(message)
		time.Sleep(400 * time.Millisecond)
	}
	channel <- "done!!!!"
}

func main() { // main goroutine

	channel := make(chan string)
	// go printMessage("hello go ")
	go printMessage("front end masters", channel)
	// go printMessage("--hello c++ ")
	response := <-channel
	fmt.Println(response)
}
