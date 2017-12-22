package main

import (
	"fmt"
)

func trocar(p1, p2 int) (segundo, primeiro int) {
	primeiro = p1
	segundo = p2
	return
}

func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, y)

	segundo, primeiro := trocar(1, 2)
	fmt.Println(segundo, primeiro)
}
