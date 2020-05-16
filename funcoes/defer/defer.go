package main

import "fmt"

func obterValorAprovado(numero int) int {
	defer fmt.Println("fim!")
	if numero >= 5000 {
		fmt.Println("Valor Alto...")
		return 5000
	}
	fmt.Println("Valor Baixo...")
	return 5000
}

func main() {
	fmt.Println(obterValorAprovado(10))
}
