package main

import (
	"fmt"
)

func rotina(c chan int) {
	c <- 1
	fmt.Println("Executou 1!")
	c <- 2
	fmt.Println("Executou 2!")
	c <- 3
	fmt.Println("Executou 3!")
	c <- 4 // Bloqueio, pois vai encher o buffer
	fmt.Println("Executou 4!")
	c <- 5
	fmt.Println("Executou 5!")
	c <- 6
}

func main() {
	ch := make(chan int, 3)
	go rotina(ch)

	fmt.Println(<-ch)
}
