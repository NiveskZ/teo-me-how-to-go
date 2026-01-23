// Faça um programa que calcule a média de uma quantidade indefinida de valores,
// usando uma função para isso.
package main

import (
	"fmt"
	"strconv"
)

func soma(values ...float64) float64 {
	res := 0.0
	for _, v := range values {
		res += v
	}

	return res
}

func media(values ...float64) float64 {
	return soma(values...) / float64(len(values))
}

func getValores() []float64 {
	values := []float64{}

	for {
		var valorStr string
		fmt.Printf("Entre com o valor: ")
		fmt.Scanf("%s\n", &valorStr)

		if valorStr == "" {
			break
		}

		valor, err := strconv.ParseFloat(valorStr, 64)
		if err != nil {
			fmt.Println("Entre com um valor valido!")
			continue
		}

		values = append(values, valor)
	}
	return values
}

func main() {
	valores := getValores()
	mediaValor := media(valores...)
	fmt.Printf("Media: %.2f", mediaValor)
}
