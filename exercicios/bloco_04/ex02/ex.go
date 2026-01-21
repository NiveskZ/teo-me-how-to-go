// Faça um programa que receba 4 alturas usando um laço de repetição e realize a soma dessas alturas
package main

import "fmt"

func main() {
	var h, soma float64
	for i := 1; i <= 4; i++ {
		fmt.Printf("Digite a %d altura \n", i)
		fmt.Scanf("%f\n", &h)
		soma += h
	}
	fmt.Printf("a soma das alturas e %.2f", soma)
}
