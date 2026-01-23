// Faça o programa de uma sorveteria, onde o usuário pode escolher:
// Tipo de sorvete: casquinha (R$1,00), cascão (R$2,50), cestinha (R$4,00)
// Sabor do sorvete: morango, creme, chocolate
// Cobertura: Caramelo (R$1,50), morango (R$1,50), chocolate (R$1,50), sem cobertura (R$0,00).
// Apresente o valor a ser pago.
package main

import (
	"fmt"
	"os"
)

func main() {
	tiposSorvete := map[string]float64{
		"casquinha": 1.00,
		"cascao":    2.50,
		"cestinha":  4.00,
	}

	saboresSorvete := map[string]float64{
		"morango":   0.00,
		"creme":     0.00,
		"chocolate": 0.00,
	}

	coberturasSorvete := map[string]float64{
		"caramelo":  1.50,
		"morango":   1.50,
		"chocolate": 1.50,
	}

	items := map[string]map[string]float64{
		"tipos":      tiposSorvete,
		"sabores":    saboresSorvete,
		"coberturas": coberturasSorvete,
	}

	var tipo, sabor, cobertura string
	fmt.Printf("Escolha um tipo [casquinha/cascao/cestinha]: ")
	fmt.Scanf("%s\n", &tipo)

	fmt.Printf("Escolha um sabor [morango/creme/chocolate]: ")
	fmt.Scanf("%s\n", &sabor)

	fmt.Printf("Escolha um tipo [caramelo/morango/chocolate]: ")
	fmt.Scanf("%s\n", &cobertura)

	total := 0.0

	if valor, ok := items["tipos"][tipo]; !ok {
		fmt.Println("Tipo invalido!")
		os.Exit(1)
	} else {
		total += valor
	}

	if valor, ok := items["sabores"][sabor]; !ok {
		fmt.Println("Sabor invalido!")
		os.Exit(1)
	} else {
		total += valor
	}

	if valor, ok := items["coberturas"][cobertura]; !ok {
		fmt.Println("Cobertura invalida!")
		os.Exit(1)
	} else {
		total += valor
	}

	fmt.Printf("\nTotal: R$%.2f\n\n", total)
}
