//Escreva uma função que recebe a e b e troque seus valores.
// a = 1; b=2 -> a=2; b=1

package main

import "fmt"

func swap(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

func swapPtrs(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	var a, b int
	fmt.Printf("Entre com a: ")
	fmt.Scanf("%d\n", &a)

	fmt.Printf("Entre com b: ")
	fmt.Scanf("%d\n", &b)

	a, b = swap(a, b)
	fmt.Printf("a = %d | b = %d\n", a, b)

	swapPtrs(&a, &b)
	fmt.Printf("a = %d | b = %d\n", a, b)
}
