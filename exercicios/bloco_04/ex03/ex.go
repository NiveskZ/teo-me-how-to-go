/*
Faça um programa que receba uma quantidade indefinida de valores correspondentes a “saldo em conta”,
mas quando o usuário apertar “enter” sem digitar valor algum, o programa para de receber valores,
e exibe a soma de todos os valores digitados anteriormente.
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var soma float64
	for true {

		var entrada string

		fmt.Printf("Digite o saldo em conta que deseja somar \n")
		fmt.Scanf("%s\n", &entrada)

		if entrada == "" {
			break
		}

		saldo_conta, erro := strconv.ParseFloat(entrada, 64)

		if erro != nil {
			fmt.Println("Entre com um valor valido!")
			continue
		}
		soma += saldo_conta
	}
	fmt.Printf("O saldo total e %.2f", soma)
}
