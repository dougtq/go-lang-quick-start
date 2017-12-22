package main

import (
	"fmt"
)

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaaraConceito(n nota) string {
	switch {
	case n.entre(9.0, 10.0):
		return "A"
	case n.entre(7.0, 8.99):
		return "B"
	case n.entre(5.0, 7.99):
		return "C"
	case n.entre(3.0, 4.99):
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaaraConceito(9.4))
	fmt.Println(notaaraConceito(7.4))
	fmt.Println(notaaraConceito(5.4))
	fmt.Println(notaaraConceito(3.4))
	fmt.Println(notaaraConceito(2.4))
}
