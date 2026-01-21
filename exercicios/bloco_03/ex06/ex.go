// Faça um programa que verifique se o item que a pessoa escolheu para comprar na loja está na lista:
// laranja, cerveja, miojo, carvão, picanha.

package main

import (
	"fmt"
	"strings"
)

func main() {
	var item string
	fmt.Println("Entre com o item:")
	fmt.Scanf("%s", &item)

	itens := "laranja, cerveja, miojo, carvão, picanha"

	if strings.Contains(itens, item) {
		fmt.Println("Esse ta tendo!")
	} else {
		fmt.Println("Esse ja nao ta tendo..")
	}
}
