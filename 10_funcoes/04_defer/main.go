package main

import "fmt"

func primeira() {
	fmt.Println("Eu sou chamada primeira")
}

func segunda() {
	fmt.Println("Eu sou chamada segunda")
}

func ultima() {
	fmt.Println("Eu sou chamada ultima")
}

func main() {
	defer ultima()
	defer segunda()
	defer primeira()

	primeira()
	segunda()
	ultima()
	fmt.Println("----------------------------------Daqui para baixo e defer----------------------------------")
}
