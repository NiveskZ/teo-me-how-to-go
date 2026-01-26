/*
Faça um programa que receba o raio de uma circunferência em centímetros.
Retorne para o usuário qual é a área e perímetro desta circunferência no seguinte formato.

Área:  x.xx
Perímetro:  y.yy
*/
package main

import (
	"fmt"
	"math"
)

func area(raio float64) (area float64) {
	area = math.Pi * math.Pow(raio, 2)
	return
}

func perimetro(raio float64) (perimetro float64) {
	perimetro = 2 * math.Pi * raio
	return
}

func main() {
	var raio float64

	fmt.Println("Me diga o raio da circunferência e vou te dizer a Área e o Perímetro!")
	fmt.Scanf("%f", &raio)

	area := area(raio)
	perimetro := perimetro(raio)
	fmt.Printf("A circunferência de raio %.2f possui:\nÁrea:\t%.2f\nPerímetro:\t%.2f", raio, area, perimetro)
}
