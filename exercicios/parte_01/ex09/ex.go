/*
Solicite ao usuário o nome de uma fruta e exiba o preço correspondente e some ao total de uma compra.

Quando o usuário inserir uma string vazia, encerre a compra e informe o valor total da compra.

Maçã: R$1,50         |		Pera: R$1,25		|		Goiaba: R$2,15

Banana: R$2,75       |		Laranja: R$0,65		|		Abacaxi: R$3,20

Uva: R$1,90          |		Limão: R$1,25		|		Jaca: R$5,80
*/
package main

import (
	"fmt"
	"strings"
)

// Utilizando struct ao inves de mapas, apenas para fixar conteudo da aula de estruturas e interface
type Estoque struct {
	Maca    float64
	Pera    float64
	Goiaba  float64
	Banana  float64
	Laranja float64
	Abacaxi float64
	Uva     float64
	Limao   float64
	Jaca    float64
}

func main() {
	estoque := Estoque{
		Maca:    1.50,
		Pera:    1.25,
		Goiaba:  2.15,
		Banana:  2.75,
		Laranja: 0.65,
		Abacaxi: 3.20,
		Uva:     1.90,
		Limao:   1.25,
		Jaca:    5.80,
	}

	var total float64
	for {
		var entrada string
		fmt.Print("Digite o nome da fruta (ou 'sair' para finalizar): ")

		_, err := fmt.Scanln(&entrada)
		if err != nil {

			fmt.Println("Erro na leitura. Encerrando.")
			return
		}

		entrada = strings.ToLower(strings.TrimSpace(entrada))
		if entrada == "sair" {
			break
		}

		switch entrada {
		case "maca", "maçã":
			total += estoque.Maca
			fmt.Printf("Maçã: R$%.2f\n", estoque.Maca)
		case "pera":
			total += estoque.Pera
			fmt.Printf("Pera: R$%.2f\n", estoque.Pera)
		case "goiaba":
			total += estoque.Goiaba
			fmt.Printf("Goiaba: R$%.2f\n", estoque.Goiaba)
		case "banana":
			total += estoque.Banana
			fmt.Printf("Banana: R$%.2f\n", estoque.Banana)
		case "laranja":
			total += estoque.Laranja
			fmt.Printf("Laranja: R$%.2f\n", estoque.Laranja)
		case "abacaxi":
			total += estoque.Abacaxi
			fmt.Printf("Abacaxi: R$%.2f\n", estoque.Abacaxi)
		case "uva":
			total += estoque.Uva
			fmt.Printf("Uva: R$%.2f\n", estoque.Uva)
		case "limao", "limão":
			total += estoque.Limao
			fmt.Printf("Limão: R$%.2f\n", estoque.Limao)
		case "jaca":
			total += estoque.Jaca
			fmt.Printf("Jaca: R$%.2f\n", estoque.Jaca)
		default:
			fmt.Println("Fruta não encontrada.")
		}

		fmt.Printf("Total parcial: R$%.2f\n\n", total)
	}

	fmt.Printf("Compra finalizada. Total da compra: R$%.2f\n", total)
}
