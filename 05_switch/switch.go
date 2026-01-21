package main

import "fmt"

func main() {
	var opcao int
	fmt.Println("Entre com uma opcao: (1,2,3,4)")
	fmt.Scanf("%d", &opcao)

	switch opcao {
	case 1:
		fmt.Println("Voce escolheu a melhor fruta!! Bananas!")
	case 2:
		fmt.Println("Eita.. Maçã! Também gosto!")
	case 3:
		fmt.Println("Top! Pera, nao e das melhores mas tambem e bom!")
	case 4:
		fmt.Println("Ahhh... Morango! esse precisa escolher bem para ser bom!")
	default:
		fmt.Println("Entre com uma opcao valida!")
	}
}
