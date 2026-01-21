package main

import "fmt"

func main() {
	nome := "Kevin"

	for i := range nome {
		fmt.Println(i, nome[i], string(nome[i]))
	}
}
