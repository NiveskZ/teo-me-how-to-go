package main

import "fmt"

type Altura float64 // em metros
type Peso float64   // em kilos

type Celsius float64
type Fahrenheit float64

func IMC(altura Altura, peso Peso) float64 {
	return float64(peso) / float64(altura*altura)
}

func ctof(c Celsius) Fahrenheit {
	return Fahrenheit((float64(c) * 9.0 / 5.0) + 32.0)
}

func main() {
	alturaTeo := Altura(1.80)
	pesoTeo := Peso(70.34)

	imc := IMC(alturaTeo, pesoTeo)
	fmt.Println("IMC teo: ", imc)

	c := Celsius(30)
	// f := Fahrenheit(15)

	fnew := ctof(c)
	fmt.Printf("%.1f C = %.1f F", c, fnew)
}
