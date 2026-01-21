// Faca um programa que exiba o dobro de um numero inserido pelo usuario
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Me de um numero")
	var x int
	fmt.Scanf("%d", &x)
	fmt.Printf("o dobro de %d e %d", x, x*2)
}
