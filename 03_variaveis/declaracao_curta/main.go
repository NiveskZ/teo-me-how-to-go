package main

import "fmt"

func main() {
	x := 45

	nome := "Kevin"

	y := 0.45

	fmt.Println(x, nome, y)

	x = 90
	fmt.Println(x)

	fmt.Printf("x = %T, nome = %T, y = %T", x, nome, y)
}
