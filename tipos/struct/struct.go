package main

import "fmt"

type produto struct {
	nome     string
	valor    float64
	desconto float64
}

func (p produto) procoCorDesconto() float64 {
	return p.valor * (1 - p.desconto)
}

func main() {
	var produto1 produto

	produto1 = produto{
		"Sabonete",
		2.00,
		0.50,
	}
	fmt.Println(produto1, produto1.procoCorDesconto())

	produto2 := produto{"Notebook", 1999.40, 200.40}
	fmt.Println(produto2, produto2.procoCorDesconto())
}
