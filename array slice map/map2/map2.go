package main

import (
	"fmt"
)

func main() {
	funcsESalarios := map[string]float64{
		"Itor":  11.23,
		"Vitor": 13.23,
		"Igor":  9.23,
	}

	funcsESalarios["Natalia"] = 41.32

	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
