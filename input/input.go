package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewReader(os.Stdin)

	fmt.Println("Digite: ")
	fmt.Print("-> ")

	texto, _ := input.ReadString('\n')

	fmt.Println(texto)

	// ler apena um caracter
	input = bufio.NewReader(os.Stdin)
	char, _, err := input.ReadRune()

	fmt.Println(char, err)

}
