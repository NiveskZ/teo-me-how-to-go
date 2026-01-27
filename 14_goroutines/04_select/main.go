package main

import (
	"fmt"
	"log"
	"time"
)

func ping(c chan string) {
	for i := 1; i <= 10; i++ {
		c <- "ping!"
		time.Sleep(time.Second * time.Duration(1))
	}
}

func pong(c chan string) {
	for i := 1; i <= 15; i++ {
		c <- "pong!"
		time.Sleep(time.Second * time.Duration(1))
	}
}

func print(c1, c2 chan string) {
	for {

		select {
		case msg1 := <-c1:
			log.Println("Canal 1:", msg1)

		case msg2 := <-c2:
			log.Println("Canal 2:", msg2)

		case <-time.After(time.Second * 3):
			fmt.Println("Tempo esgotado")
		}

	}
}

func main() {
	// var canal chan string = make(chan string)
	canal := make(chan string)
	canal2 := make(chan string)

	go ping(canal)
	go pong(canal2)
	go print(canal, canal2)

	var ok string
	fmt.Scanf("%s", &ok)

}
