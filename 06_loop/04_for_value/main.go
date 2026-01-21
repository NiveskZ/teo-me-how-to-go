package main

import "fmt"

func main() {
	nome := "Kevin"

	for _, v := range nome {
		fmt.Println(v, string(v))
	}
}
