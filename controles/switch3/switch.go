package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "funcão"
	case bool:
		return "Boolean"
	case time.Time:
		return "Tempo"
	default:
		return "Não sei"

	}
}

func main() {
	fmt.Println(tipo(2))
	fmt.Println(tipo(2.3))
	fmt.Println(tipo("2.3"))
	fmt.Println(tipo(func() {

	}))
	fmt.Println(tipo(true))
	fmt.Println(tipo(time.Now()))
}
