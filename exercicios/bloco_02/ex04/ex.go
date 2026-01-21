// Faca um programa que recebe um numero inteiro e calcula sua raiz quadrada
package main

import (
	"fmt"
	"math"
)

func main() {
	var x int

	fmt.Println("Me diga um numero inteiro")
	fmt.Scanf("%d", &x)
	y := math.Sqrt(float64(x))
	fmt.Printf("A raiz quadrada de %d é %.2f\n", x, y)
}
