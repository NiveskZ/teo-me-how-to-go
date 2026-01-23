package main

import "fmt"

func main() {

	idades := make(map[string]int)

	idades["teo"] = 33
	fmt.Println(idades)

	alturas := map[string]float64{}
	alturas["teo"] = 1.82

	fmt.Println(alturas)

	alturaTeo := alturas["teo"]
	fmt.Println("Altura Teo:", alturaTeo)

	alturaNah, ok := alturas["nah"]

	if ok {
		fmt.Println("Altura Nah:", alturaNah)
	} else {
		fmt.Println("Nao encontrei!")
	}

	if valor, ok := alturas["kevin"]; ok {
		fmt.Println("A altura do Kevin e:", valor)
	} else {
		fmt.Println("Nao encontrei!")
	}
}
