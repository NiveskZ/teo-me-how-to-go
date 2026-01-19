package main

import "fmt"

func main() {
	fmt.Println("Com aspas duplas")

	fmt.Println(`Com crase
Daí
	Posso
		Escrever
	Assim
Se eu quiser!`)

	fmt.Println("Meu nome é Kevin e tem", len("Kevin"), "letras")

	fmt.Println(string("Kevin"[0]))
	fmt.Println(string("Kevin"[1]))
	fmt.Println(string("Kevin"[2]))
}
