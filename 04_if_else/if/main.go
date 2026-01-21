package main

import "fmt"

func main() {

	var idade int
	fmt.Println("Entre com a sua idade:")
	fmt.Scanf("%d", &idade)

	if idade >= 18 {
		fmt.Println("Perfeito!")
		fmt.Println("Beba a vontade! Mas com moderacao...")
	}

	if idade <= 17 {
		fmt.Println("Vishhhhhh")
		fmt.Println("Melhor nao beber!")
	}
}
