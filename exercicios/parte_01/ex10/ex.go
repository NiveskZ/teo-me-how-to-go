// Escreva um programa que exiba os números de 1 a 100.
// Caso o número seja divisível por 3, exiba “Fizz” no seu lugar, e para múltiplos de 5 exiba “Buzz”.
// Caso seja divisível por ambos, exiba “FizzBuzz”.

package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

// Metodo 2
