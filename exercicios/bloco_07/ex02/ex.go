// Faça um programa que informe se uma número é par ou ímpar
// com uma função que se chama “EhPar”, retornando um booleano.
package main

import "fmt"

func ehPar(num int) bool {
	return num%2 == 0
}

func main() {
	var numero int
	fmt.Printf("Entre com um numero: ")
	fmt.Scanf("%d", &numero)

	if ehPar(numero) {
		fmt.Println("Par!")
	} else {
		fmt.Println("Impar!")
	}
}
