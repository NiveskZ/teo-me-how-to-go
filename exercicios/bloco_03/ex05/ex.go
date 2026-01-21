// Faça um programa que verifique se a pessoa pertence à família “calvo” ou “silva”.
package main

import (
	"fmt"
	"strings"
)

func main() {
	var nome string
	fmt.Println("Digite seu nome completo: ")
	fmt.Scanf("%s", &nome)

	if strings.Contains(nome, "Calvo") || strings.Contains(nome, "Silva") {
		fmt.Println("Claramente voce eh Calvo! Ou, talvez, um Silva!")
	} else {
		fmt.Println("Voce nao eh careca e Nem Silva")
	}
}
