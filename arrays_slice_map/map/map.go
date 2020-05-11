package main

import "fmt"

func main() {
	// var aprovados map[int]string // chave int e valor string
	// map deve ser inicializado
	aprovados := make(map[int]string)
	aprovados[651561156] = "Maria"
	aprovados[984651682] = "Zé"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Println(cpf, nome)
	}

	delete(aprovados, 984651682)
}
