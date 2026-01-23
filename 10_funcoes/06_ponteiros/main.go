package main

import "fmt"

func addOne(a int) int {
	fmt.Println("Endereco de a: ", &a)
	return a + 1
}

func addOnePt(a *int) {
	fmt.Println("Endereco do ponteiro a: ", a)
	fmt.Println("valor do ponteiro de a: ", *a)
	*a++
}

func main() {
	numero := 1
	fmt.Println("Endereco de numero: ", &numero)
	valor := addOne(numero) // isso gera uma copia, nao e o mesmo objeto

	addOnePt(&numero)
	fmt.Println("Numero: ", numero)
	fmt.Println("Valor", valor)
}
