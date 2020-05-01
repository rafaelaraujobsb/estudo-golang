package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	// fmt.Println(x / y) ERRO
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado, ao converte um inteiro para string, ele chama o valor da tabela ASCII
	fmt.Println("Teste " + string(97))

	// int para string
	fmt.Println("Teste " + strconv.Itoa(97))

	// string para int
	num, _ := strconv.Atoi("123") // o segundo valor é um erro, caso ocorra
	fmt.Println(reflect.TypeOf(num))

	// string para bool
	b, _ := strconv.ParseBool("true") // o segundo valor é um erro, caso ocorra
	fmt.Println(reflect.TypeOf(b))
}
