/*
Faça um programa que vende uma garrafa de água:
- Se o cliente escolher água mineral natural, será cobrado R$1,50
- Se o cliente escolher água mineral com gás, será cobrado R$2,50
*/
package main

import "fmt"

func main() {

	// 1 - Sem gas = 1,50
	// 2 - Com gas = 2,50
	var x int

	fmt.Println("Agua com ou sem gas chefe?")
	fmt.Println("Sem Gas - R$ 1,50 (1)")
	fmt.Println("Com Gas - R$ 2,50 (2)")
	fmt.Scanf("%d", &x)

	if x == 1 {
		fmt.Println("Me deve R$ 1,50")
	} else if x == 2 {
		fmt.Println("Me deve R$ 2,50")
	} else {
		fmt.Println("To sem chefe!")
	}
}
