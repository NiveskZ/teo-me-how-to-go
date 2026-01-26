// Faça um programa que receba dois valores A e B.
// Faça a soma desses dois valores e retorne o resultado:
// Soma:  x.xx
package main

import "fmt"

func soma(a, b float64) (res float64) {
	res = a + b
	return
}

func main() {
	var a, b float64

	fmt.Println("Irei somar dois valores para voce, digite o primeiro valor: ")
	fmt.Scanf("%f\n", &a)
	fmt.Println("Agora o segundo valor: ")
	fmt.Scanf("%f\n", &b)

	res := soma(a, b)

	fmt.Printf("Soma:\t%.2f", res)
}
