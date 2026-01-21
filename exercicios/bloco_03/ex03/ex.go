/*
Faça o programa de uma sorveteria, onde o usuário pode escolher:

	Tipo de sorvete: casquinha (R$1,00), cascão (R$2,50), cestinha (R$4,00)
	Sabor do sorvete: morango, creme, chocolate
	Cobertura: Caramelo (R$1,50), morango (R$1,50), chocolate (R$1,50), sem cobertura (R$0,00).
	Apresente o valor a ser pago
*/
package main

import "fmt"

func main() {
	var x, sabor, cober string
	fmt.Println("Qual tipo de sorvete voce quer? casquinha, cascao, cestinha")
	fmt.Scanf("%s", &x)

	switch x {
	case "casquinha":
		fmt.Println("Qual Sabor? Morango, Creme, Chocolate")
		fmt.Scanf("%s\n", &sabor)

		fmt.Println("Cobertura? Caramelo, morango ou chocolate")
		fmt.Scanf("%s", &cober)

		if cober == "nao" {
			val := 1.00
			fmt.Printf("Deu %f", val)
		} else {
			val := 1.00 + 1.50
			fmt.Printf("Deu %f", val)
		}

	case "cascao":
		fmt.Println("Qual Sabor? Morango, Creme, Chocolate")
		fmt.Scanf("%s\n", &sabor)

		fmt.Println("Cobertura? Caramelo, morango ou chocolate")
		fmt.Scanf("%s", &cober)

		if cober == "nao" {
			val := 2.50
			fmt.Printf("Deu %f", val)
		} else {
			val := 2.50 + 1.50
			fmt.Printf("Deu %f", val)
		}
	case "cestinha":
		fmt.Println("Qual Sabor? Morango, Creme, Chocolate")
		fmt.Scanf("%s\n", &sabor)

		fmt.Println("Cobertura? Caramelo, morango ou chocolate")
		fmt.Scanf("%s", &cober)

		if cober == "nao" {
			val := 4.00
			fmt.Printf("Deu %f", val)
		} else {
			val := 4.00 + 1.50
			fmt.Printf("Deu %f", val)
		}
	default:
		fmt.Println("Nao tem, ou voce digitou errado! garanta letras minusculas e sem acentos: casquinha, cascao, cestinha")
	}

}
