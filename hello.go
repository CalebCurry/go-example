package main

import (
	"fmt"
	"time"
)

func greet(name string, channel chan string) {
	channel <- fmt.Sprintln("hello", name)
}

func greetByLetter(name string, channel chan string) {
	time.Sleep(3 * time.Second)
	channel <- fmt.Sprintln("hello", name)
}

func main() {
	channel := make(chan string)
	go greet("John", channel)
	go greetByLetter("Sal", channel)
	for i := 0; i < 2; i++ {
		select {
		case response := <-channel:
			fmt.Println("received:", response)
		case <-time.After(4 * time.Second):
			fmt.Println("Timeout occured")
		}
	}
	// message := <-channel
	// fmt.Println(message)
	// message = <-channel
	// fmt.Println(message)
}
