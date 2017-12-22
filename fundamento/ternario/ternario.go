package main

import (
	"fmt"
)

func obterREsultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(obterREsultado(6, 2))
}
