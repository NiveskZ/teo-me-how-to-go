// Escreva um programa que receba uma lista de números do usuário
// e conte quantas vezes um número específico aparece na lista.
// Solicite ao usuário um número e exiba a contagem.
package main

import "fmt"

func main() {
	values := map[int]int{}

	for {
		var v int
		fmt.Println("Entre com um numero para ser registrado")
		fmt.Scanf("%d\n", &v)

		if v == 0 {
			break
		}
		values[v]++
	}

	busca := 0
	fmt.Printf("Entre com um numero para descobrir quantas vezes foi digitado: ")
	fmt.Scanf("%d", &busca)

	fmt.Printf("O numero %d apareceu %d vezes.", busca, values[busca])
}
