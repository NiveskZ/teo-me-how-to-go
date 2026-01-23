package main

import "fmt"

func soma(values ...int) int {

	total := 0
	for _, v := range values {
		total += v
	}
	return total

}

func main() {
	var a, b, c, d int

	a = 10
	b = 20
	c = 40
	d = 30

	total := soma(a, b, c, d)
	fmt.Println(total)

	valor := []int{1, 2, 3, 4, 23, 48, 32, 42, 421, 54}
	total = soma(valor...)
	fmt.Println(total)
}
