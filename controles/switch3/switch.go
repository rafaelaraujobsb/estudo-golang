package main

import "fmt"

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case func():
		return "função"
	default:
		return "outro"
	}
}

func main() {
	fmt.Println(tipo("ABS"))
	fmt.Println(tipo(1))
}
