package main

import "fmt"

func main() {
	i := 1

	//  declaração de um ponteiro
	var p *int = nil

	p = &i

	fmt.Println("Conteúdo: ", *p)
	fmt.Println("Endreço p: ", p)
	fmt.Println("Endreço i: ", &i)
}
