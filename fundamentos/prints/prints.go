package main

import "fmt"

func main() {
	fmt.Print("Teste")
	fmt.Print("!")
	fmt.Println("<br>")
	fmt.Println("new")

	x := 3.141516
	fmt.Printf("O valor é %.2f\n", x)
	// fmt.Println("O valor é " + x)
	xs := fmt.Sprint(x)
	fmt.Println("O valor é " + xs)
}
