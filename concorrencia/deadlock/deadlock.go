package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1
	fmt.Println("Depois que ser lida")
}

func main() {
	c := make(chan int) // canal sem buffer
	go rotina(c)
	fmt.Println(<-c)
	fmt.Println("Foi lido")
	fmt.Println(<-c)
	fmt.Println("Fim")
}
