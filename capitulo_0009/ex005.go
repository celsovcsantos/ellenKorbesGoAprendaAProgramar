package capitulo_0009

import (
	"fmt"
)

func Ex005() {

	fmt.Println("Exercício 005 - OK")

	fmt.Println("1ª questão")
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Printf("%v\n\n", slice)
	slice = append(slice[:3], slice[6:]...)
	fmt.Printf("%v\n\n", slice)

	fmt.Println("")
	fmt.Println("")
}
