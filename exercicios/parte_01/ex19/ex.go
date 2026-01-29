// Escreva um programa que solicite ao usuário frases.
// Para parar de solicitar frases, ele pode apenas apertar o “enter”.
// Seu programa deve apresentar cada frase e quantas vezes ela foi repetida.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	frases := map[string]int{}
	reader := bufio.NewReader(os.Stdin)
	for {
		r, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			break
		}

		r = strings.TrimSuffix(r, "\n")
		r = strings.TrimSuffix(r, "\r")

		if r == "" {
			break
		}

		frases[r]++
	}
	for f, q := range frases {
		fmt.Printf("%s: %d\n", f, q)
	}
}
