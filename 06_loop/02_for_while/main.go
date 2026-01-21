package main

import "fmt"

func main() {
	i := 14
	for true {
		if i%13 == 0 {
			fmt.Println("while", i)
			break
		}
		i++
	}
}
