package main

import (
	"fmt"
)

func imprimir(dado map[string]map[string]float64) {
	for letra, funcs := range dado {
		for nome, valor := range funcs {
			fmt.Println(letra, nome, valor)
		}
		fmt.Println("-------")
	}
}

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"C": {
			"Carol": 11.2,
			"Casia": 32.4,
		},
		"I": {
			"Itor": 400.0,
			"Igor": 40.0,
		},
	}
	imprimir(funcsPorLetra)
	delete(funcsPorLetra, "C")
	fmt.Println("Depois de deletar todos com C")
	imprimir(funcsPorLetra)
}
