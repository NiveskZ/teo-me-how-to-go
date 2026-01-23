// Faça um programa que receba 6 temperaturas.
// Remova a 1a e a última para calcular a média.
// Se a média for acima de 30 graus, exiba que houve um aumento da temperatura.
package main

import "fmt"

func main() {
	var temperaturas [6]float64

	for i := 1; i <= 6; i++ {

		var temp float64
		fmt.Printf("Entre com a %d temperatura: ", i)
		fmt.Scanf("%f\n", &temp)
		temperaturas[i-1] = temp
	}

	temps := temperaturas[1 : len(temperaturas)-1]

	fmt.Println(temps)

	var total float64

	for _, temp := range temps {
		total += temp
	}

	media := total / float64(len(temps))

	if media > 30 {
		fmt.Println("Esta ficando mais quente! A media esta em: ", media)
	} else {
		fmt.Println("Media: ", media)
	}
}
