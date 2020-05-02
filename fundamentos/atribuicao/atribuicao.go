package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 // inferência de tipo
	i += 3

	fmt.Println(i)

	x, y := 2, 3

	fmt.Println(x, y)

	x, y = y, x
}
