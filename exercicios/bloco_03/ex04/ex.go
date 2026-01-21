// Faça um programa que verifique se a pessoa pertence à família “calvo”.
package main

import (
	"fmt"
	"strings"
)

func main() {
	var nome string
	fmt.Println("Digite seu nome completo: ")
	fmt.Scanf("%s", &nome)

	if strings.Contains(nome, "Calvo") {
		fmt.Println("Claramente voce eh Calvo!")
	} else {
		fmt.Println("Voce nao eh careca! Mas sempre da para raspar a cabeca...")
	}
}
