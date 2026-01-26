// Escreva um programa que receba o nome e a idade de uma pessoa
// Depois exiba a mensagem:
//
//	“Olá fulano, bom saber que você tem x anos. Boas vindas!”
package main

import "fmt"

func nome_idade() (nome string, idade int) {

	fmt.Println("Oi! Qual seu nome?")
	fmt.Scanf("%s\n", &nome)
	fmt.Println("E sua idade?")
	fmt.Scanf("%d", &idade)

	return
}

func main() {
	nome, idade := nome_idade()
	fmt.Printf("Olá %s, bom saber que você tem %d anos. Boas vindas!", nome, idade)

}
