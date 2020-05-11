package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"Z": {
			"Zé": 15186.56,
		},
		"M": {
			"Maria":   4684.44,
			"Mariana": 15641.78,
		},
	}

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra)
		for nome := range funcs {
			fmt.Printf("  >> %s\n", nome)
		}
	}
}
