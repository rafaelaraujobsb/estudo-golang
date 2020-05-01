package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 100)
	fmt.Println("Literal inteiro é", reflect.TypeOf(100))

	// sem sinal (só positivo) ... uint8, uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal ... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo é", i1)
	fmt.Println("O tipo é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))

	// números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) // float64

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// String, somente aspas duplas
	s1 := "Olá"
	fmt.Println("Tamanho", len(s1)) // o acento conta no tamanho

	// string multipla
	s2 := ` Olá
	teste
	`
	fmt.Println("Tamanho", len(s2)) // o acento conta no tamanho

	// o char não existe no go
	char := 'a'
	fmt.Println("Tipo", reflect.TypeOf(char))
	fmt.Println(char)
}
