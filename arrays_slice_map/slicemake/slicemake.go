package main

import "fmt"

func main() {
	s := make([]int, 10) // cria um slice com um array interno
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20) // 10 elementos e o array interno terá 20 posições, ou seja, a capacidade desse slice
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0) // o slice continua o mesmo, ou seja, mesmo endereço de memória
	ptr := &s
	fmt.Printf("%p\n", ptr)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1)
	ptr = &s
	fmt.Printf("%p\n", ptr)
	fmt.Println(s, len(s), cap(s)) // ele vai crescer de acordo com a sua capacidade, nesse caso cresceu para 40
}
