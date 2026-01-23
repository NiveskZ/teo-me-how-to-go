package main

import (
	"fmt"
	"strconv"
)

func validaEntrada() float64 {

	defer func() {
		txt := recover()
		fmt.Println(txt)
	}()

	var input string
	fmt.Println("Entre com um float: ")
	fmt.Scanf("%s\n", &input)

	num, err := strconv.ParseFloat(input, 64)

	if err != nil {
		panic("Entrada invalida")
	}

	return num
}

func main() {

	num := validaEntrada()
	fmt.Println(num)
}
