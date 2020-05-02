package main

import (
	"fmt"
	"math/rand"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i > 5 { // atribuição na operação
		fmt.Println("Ganhou!")
	} else {
		fmt.Println("Perdeu!")
	}

	fmt.Println(i) // não está disponível fora do escopo
}
