package main

import (
	"fmt"
	"strconv"
)

func main() {
	var notas []float64

	for {
		var inputTxt string
		fmt.Printf("Entre com a nota: ")
		fmt.Scanf("%s\n", &inputTxt)

		if inputTxt == "" {
			break
		}

		nota, err := strconv.ParseFloat(inputTxt, 64)
		if err != nil {
			fmt.Println("Entre com uma nota valida")
		}

		notas = append(notas, nota)
	}

	fmt.Println(notas)
}
