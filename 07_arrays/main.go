package main

import "fmt"

func main() {

	var x [20]float64
	fmt.Println((x))

	var y = [20]float64{1, 32, 42, 9.99, 10}
	fmt.Println("y =", y)
	fmt.Println("Primeiro elemento de y = y[0] = ", y[0])
	fmt.Println("Dois Primeiros elementos de y = y[0:2] = ", y[0:2])

	y[18] = 33
	y[19] = 44
	fmt.Println("Dois ultimos elementos de y = y[18:] = ", y[18:])
	fmt.Println("Dois ultimos elementos de y = y[len(y)-2:] = ", y[len(y)-2:])

	var total float64
	for i, v := range y {
		fmt.Println("Indice=", i, "| Valor =", v)
		total += v
	}

	fmt.Println("Soma de y =", total)

	t := [10]string{"Teo", "Calvo"}
	fmt.Println(t)
}
