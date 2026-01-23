// Faça um programa que receba 4 notas, calcule a média, mínimo e máximo dessas notas.
package main

import (
	"fmt"
	"slices"
	"strconv"
)

func main() {
	var notas [4]float64

	for i := range notas {
		var inputTxt string

		fmt.Println("Entre com a nota: ")
		fmt.Scanf("%s\n", &inputTxt)

		nota, err := strconv.ParseFloat(inputTxt, 64)

		if err != nil {
			fmt.Println("Passe um valor valido")
		}

		notas[i] = nota
	}
	var total float64

	for _, nota := range notas {
		total += nota
	}
	min := slices.Min(notas[:])
	max := slices.Max(notas[:])
	media := total / float64(len(notas))

	fmt.Printf("Media = %.2f, Minimo = %.2f, Maximo = %.2f", media, min, max)
}
