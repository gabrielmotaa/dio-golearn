package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for {
		c <- "ping"
	}
}

func pong(c chan string) {
	for {
		c <- "pong"
	}
}

func printMsg(c chan string) {
	for {
		select {
		case msg := <- c:
			fmt.Println(msg)
		}
	}
}

func main() {
	c := make(chan string)

	go ping(c)
	go printMsg(c)
	go pong(c)

	// Evitar o loop infinito
	<- time.After(time.Second)

}
