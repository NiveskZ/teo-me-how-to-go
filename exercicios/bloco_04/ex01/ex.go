// Faça um programa que conte quantas vezes a letra “a” aparece em uma palavra
package main

import (
	"fmt"
	"strings"
)

func main() {
	var palavra string
	fmt.Println("Digite uma palavra!")
	fmt.Scanf("%s", &palavra)

	palavra = strings.ToLower(palavra)

	j := 0
	for _, v := range palavra {
		if string(v) == "a" {
			j++
		}

	}

	fmt.Printf("A palavra %s possui %d letras 'a'", palavra, j)
}
