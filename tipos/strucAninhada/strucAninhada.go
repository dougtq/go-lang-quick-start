package main

import (
	"fmt"
)

type item struct {
	produtoID int
	qtd       int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtd)
	}
	return total
}

func main() {
	pedido1 := pedido{
		userID: 1,
		itens: []item{
			item{produtoID: 1, qtd: 1, preco: 22.00},
			item{2, 3, 2.00},
			item{3, 2, 17.00},
		},
	}

	fmt.Printf("O pedido do Usuario %d\n", pedido1.userID)
	fmt.Println("N | produto | quantidade | valor")
	fmt.Println("-------------------------------------")
	for index, iem := range pedido1.itens {
		fmt.Printf("%d | %d       | %d          | %f \n", index+1, iem.produtoID, iem.qtd, iem.preco)
	}
	fmt.Println("-------------------------------------")
	fmt.Printf("Total                      R$%.2f\n", pedido1.valorTotal())
}
