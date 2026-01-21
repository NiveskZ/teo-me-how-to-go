/*
Faça um programa que vende uma garrafa de água:
- Se o cliente escolher água mineral natural, será cobrado R$1,50
- Se o cliente escolher água mineral com gás, será cobrado R$2,50

Altere o programa anterior para considerar a quantidade de garrafas de água
*/
package main

import "fmt"

func main() {

	// 1 - Sem gas = 1,50
	// 2 - Com gas = 2,50
	var x int
	var qtd float64

	fmt.Println("Agua com ou sem gas chefe?")
	fmt.Println("Sem Gas - R$ 1,50 (1)")
	fmt.Println("Com Gas - R$ 2,50 (2)")
	fmt.Println("Me fala a quantidade ae tambem")
	fmt.Scanf("%d %f", &x, &qtd)

	if x == 1 {
		val := qtd * 1.50
		fmt.Printf("Me deve R$ %f", val)
	} else if x == 2 {
		val := qtd * 2.50
		fmt.Printf("Me deve R$ %f", val)
	} else {
		fmt.Println("To sem chefe!")
	}
}
