/*
Escreva um programa que solicite ao usuário seu nome e depois seu sobrenome.

Exiba ambos depois utilizando apenas uma variável.
*/

package main

import "fmt"

func main() {
	var nome, sobrenome string
	fmt.Println("Qual seu primeiro nome?")
	fmt.Scanf("%s\n", &nome)
	fmt.Println("Qual seu sobrenome nome?")
	fmt.Scanf("%s\n", &sobrenome)

	nome_sobrenome := nome + " " + sobrenome
	fmt.Println(nome_sobrenome)
}
