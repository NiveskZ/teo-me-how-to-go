// Escreva um programa com a função “quadrado”,
// que recebe um número inteiro e eleva ele ao quadrado, retornando o resultado.

package main

import "fmt"

func quadrado(x int) int {
	return x * x
}

func main() {
	var x int
	fmt.Printf("Entre com um numero inteiro: ")
	fmt.Scanf("%d\n", &x)

	q := quadrado(x)

	fmt.Printf("%d x %d = %d", x, x, q)
}
