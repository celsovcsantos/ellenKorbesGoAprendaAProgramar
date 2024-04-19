package capitulo_0009

import (
	"fmt"
)

func Ex004() {

	fmt.Println("Exercício 004 - OK")

	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println("1ª questão")
	slice = append(slice, 52)
	fmt.Printf("%v\n\n", slice)

	fmt.Println("2ª questão")
	slice = append(slice, 53, 54, 55)
	fmt.Printf("%v\n\n", slice)

	fmt.Println("")
	fmt.Println("")
}
