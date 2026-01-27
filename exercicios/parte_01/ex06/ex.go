/*
Faça um programa que receba um número em segundos, converta esse número para horas, minuto e segundos.

Exemplos:

Entrada: 556
Saída: 0:9:16

Entrada: 140153
Saída: 38:55:53
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	var s int
	fmt.Println("Digite os Segundos: ")
	fmt.Scanf("%d\n", &s)

	d := time.Duration(s) * time.Second

	horas := d / time.Hour
	d %= time.Hour
	minutos := d / time.Minute
	segundos := d % time.Minute / time.Second

	fmt.Printf("Duracao: %d:%d:%d\n", horas, minutos, segundos)

}
