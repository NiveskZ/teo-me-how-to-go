package main

import "fmt"

func soma(a int, b int) int {
	res := a + b
	return res
}

func media(a, b int) (res float64, err error) {

	total := float64(soma(a, b))
	res = total / 2

	err = nil

	return
}

func main() {
	a, b := 10, 20
	fmt.Println(a, b)

	total := soma(b, a)
	fmt.Println("Total:", total)

	resMedia, erro := media(a, b)
	if erro != nil {
		fmt.Println("Erro!")
	}
	fmt.Println("Media:", resMedia)
}
