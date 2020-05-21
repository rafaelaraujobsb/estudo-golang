package main

import "fmt"

type carro struct {
	nome            string
	valocidadeAtual int
}

type ferrari struct {
	carro       // campos anonimos
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.turboLigado = false
	fmt.Printf("Carro: %s\nTurbo ligado? %v\n", f.nome, f.turboLigado)
	fmt.Println(f)
}
