package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch { // verifica se algum case é verdadeiro
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde!")
	}
}
