package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 { // parece com o while
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d", i)
	}

	for {
		// laço infinito
		fmt.Println("Pra sempre...")
		time.Sleep(time.Second)
	}
}
