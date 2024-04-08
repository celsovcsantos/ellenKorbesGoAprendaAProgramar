package capitulo_0005

import (
	"fmt"
)

const (
	_ = 2024 + iota
	a_1
	a_2
	a_3
	a_4
)

func Ex006() {
	fmt.Println("Exercício 006 - OK")
	fmt.Print("1ª questão: ")
	fmt.Printf("%v\t%v\t%v\t%v\n", a_1, a_2, a_3, a_4)
	fmt.Printf("")
	fmt.Printf("")
}
