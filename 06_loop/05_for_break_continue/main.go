package main

import "fmt"

func main() {
	nome := "Kevin Richardson"

	for _, v := range nome {

		letra := string(v)
		if letra == "h" {
			break
		}

		if letra == "R" {
			continue
		}

		fmt.Println(letra)
	}
}
