/*
	Faca um programa que de bom dia,

pergunta o nome da pessoa e
responde que e um prazer conhecer ela, citando o nome da pessoa
*/
package main

import "fmt"

func main() {
	fmt.Println("Bom Dia! Qual seu nome?")

	var nome string
	fmt.Scanf("%s", &nome)

	fmt.Printf("Um prazer te conhecer %s", nome)
}
