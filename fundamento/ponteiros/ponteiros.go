package main

import (
	"fmt"
)

func main() {
	i := 1

	var p *int = nil

	p = &i // pega endereço da var

	*p++

	fmt.Println(p, *p, i, &i)
}
