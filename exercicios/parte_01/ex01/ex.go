// Escreva um programa que receba o nome de uma pessoa e faça uma saudação.
package main

import "fmt"

func saudacao(nome string) {
	fmt.Printf("Saudação %s! Espero que esteja tudo bem!", nome)
}

func main() {
	var nome string
	fmt.Println("Oi! Qual seu nome?")
	fmt.Scanf("%s", &nome)
	saudacao(nome)
}
