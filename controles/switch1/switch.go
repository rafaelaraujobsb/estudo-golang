package main

import "fmt"

func imprimirResultado(n float64) string {
	var nota = int(n)

	switch nota {
	case 10:
		fallthrough // com isso ele continua verificando os outros cases
	case 9:
		return "A"
	case 8, 7:
		return "B"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(imprimirResultado(7))
}
