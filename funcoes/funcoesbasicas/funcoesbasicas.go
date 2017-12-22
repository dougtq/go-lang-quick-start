package main

import (
	"fmt"
)

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "Retornou 1", "Retornou 2"
}

func main() {
	f1()
	f2("asd", "abc")
	fmt.Println(f3())
	fmt.Println(f4("asd", "abc"))
	fmt.Println(f5())

	r1, _ := f5()
	fmt.Println(r1)

	_, r2 := f5()
	fmt.Println(r2)

}
