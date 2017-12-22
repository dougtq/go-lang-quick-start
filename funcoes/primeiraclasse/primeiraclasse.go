package main

import (
	"fmt"
)

var soma = func(a, b int) int {
	return a + b
}

func main() {
	resultado := soma(1, 2)

	fmt.Println(resultado)

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(3, 1))

}
