package main

import "fmt"

func main() {
	fmt.Println("-------- TABELA VERDADE ----------")
	fmt.Println(" p | q | p or q | p and q | not p")
	fmt.Println(" V | V |", true || true, "|", true && true, "|", !true)
	fmt.Println(" V | F |", true || false, "|", true && false, "|", !true)
	fmt.Println(" F | V |", false || true, "|", false && true, "|", !false)
	fmt.Println(" F | F |", false || false, "|", false && false, "|", !false)
}
