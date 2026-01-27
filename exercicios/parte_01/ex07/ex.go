/*
Faça um programa que receba um número.
Verifique se o número informado é par ou ímpar.

Exiba o resultado da seguinte maneira:

	O número x é impar

ou

	O número x é par
*/
package main

import "fmt"

func main() {
	var x int

	n, err := fmt.Scanf("%d", &x)

	if n != 1 || err != nil {
		fmt.Println("Entrada inválida: digite um inteiro (sem ponto).")
		return

	} else if x%2 == 0 {
		fmt.Printf("O número %d é par", x)
	} else {
		fmt.Printf("O número %d é impar", x)
	}

}
