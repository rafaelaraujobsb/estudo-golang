package main

import "fmt"

func main() {
	ch := make(chan int, 1) // o segundo parâmetro é o buffer
	ch <- 1                 // enviando dados para o canal
	<-ch                    // recebendo dados do canal

	ch <- 2
	fmt.Println(<-ch)
}
