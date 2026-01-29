// Faça um programa que receba um número e exiba seu fatorial.
// Use função para isso.
package main

import "fmt"

func fatorial(a int) int {
	result := 1
	for i := a; i >= 1; i-- {
		result *= i
	}

	return result
}

func main() {
	var x int
	fmt.Println("Entre com um numero: ")
	fmt.Scanf("%d\n", &x)

	res := fatorial(x)
	fmt.Printf("%d! = %d", x, res)
}
