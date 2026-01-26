// Faça um programa que receba dois valores A e B.
// Faça a potência desses dois valores e retorne o resultado:
// a ^ b = z
package main

import "fmt"

// Criando a funcao ao inves de usar math.pow()
func pow(a float64, b int) (res float64) {
	if b == 0 {
		return 1
	}

	if b < 0 {
		return 1 / pow(a, -b)
	}

	res = 1.0
	base := a
	for b > 0 {
		if b&1 == 1 {
			res *= base
		}
		base *= base
		b >>= 1
	}
	return
}

func main() {
	var a float64
	var b int
	fmt.Println("Digite o valor de A: ")
	fmt.Scanf("%f\n", &a)
	fmt.Println("Digite o valor de B: ")
	fmt.Scanf("%d\n", &b)

	res := pow(a, b)
	fmt.Printf("%.2f", res)
}
