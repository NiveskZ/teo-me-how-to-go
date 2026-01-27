package main

import (
	"fmt"
	"log"
	"time"
)

func dizAlgo(txt string) {
	for {
		log.Println(txt)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	go dizAlgo("Ola")
	go dizAlgo("Tchau")

	var ok string
	fmt.Scanf("%d", &ok)
	log.Println("Fim do programa!!")
}
