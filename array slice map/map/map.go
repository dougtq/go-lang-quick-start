package main

import (
	"fmt"
)

func main() {
	// var aprovados map[int]string
	aprovados := make(map[int]string)

	aprovados[123456789] = "Itor"
	aprovados[123456783] = "Vitor"
	aprovados[123456782] = "Igor"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF :%d)\n", nome, cpf)
	}

	fmt.Println(aprovados[123456782])
	delete(aprovados, 123456782)
	fmt.Println(aprovados)
}
