package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	for i := 0; i <= 10; i++ {
		if i&2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Impar")
		}
	}

	for {
		// Laco infinito
		fmt.Println("Para Sempre... \n Para parar aperte CTRL + C")
		time.Sleep(time.Second * 30)
	}

	// for {i <- 0 to 60; j <- 0 to 10}
}
