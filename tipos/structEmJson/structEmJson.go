package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	p1 := produto{1, "Notebook", 1989.29, []string{"Promo", "Eletronico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	var p2 produto
	jsonString := `{"id":2,"nome":"Caneta","preco":1.50,"tags":["Papelaria","Importate"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[0])
}
