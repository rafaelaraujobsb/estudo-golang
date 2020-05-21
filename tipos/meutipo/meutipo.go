package main

import "fmt"

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	var resp string
	if n.entre(9.0, 10.0) {
		resp = "A"
	} else if n.entre(7.0, 8.99) {
		resp = "B"
	} else {
		resp = "C"
	}
	return resp
}

func main() {
	fmt.Println(notaParaConceito(8.0))
}
