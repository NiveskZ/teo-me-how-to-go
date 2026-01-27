package main

import (
	"log"
	"time"
)

func ping(c chan string) {
	for {
		c <- "ping!"
	}
}

func pong(c chan string) {
	for {
		c <- "pong!"
		log.Println("Enviado PONG!")
	}
}

func main() {
	canal := make(chan string, 100)

	go ping(canal)
	go pong(canal)

	for {
		msg := <-canal
		log.Println("Saida:", msg)
		time.Sleep(time.Second)
	}
}
