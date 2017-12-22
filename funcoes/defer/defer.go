package main

import (
	"fmt"
)

func obterValorAprovado(n int) int {
	defer fmt.Println("Fim")
	if n > 5000 {
		fmt.Println("Valor Alto...")
		return 5000
	}
	fmt.Println("Valor Baixo....")
	return n
}

func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
}
