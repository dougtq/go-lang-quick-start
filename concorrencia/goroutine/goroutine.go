package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Itor", "Olá Natlaia", 4)
	// fale("Natalia", "Sai seu ", 1)

	// go fale("Itor", "Ei.......", 500)
	// go fale("Natalia", "kkkkkk", 500)

	// time.Sleep(time.Second * 5)

	// fmt.Println("Fim")

	go fale("Itor", "Entendi", 10)
	fale("Natalia", "A Blz", 5)
}
