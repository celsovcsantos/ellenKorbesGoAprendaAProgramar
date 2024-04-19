package capitulo_0009

import (
	"fmt"
)

func Ex003() {

	fmt.Println("Exercício 003 - OK")

	slice := []int{}

	for i := 0; i < 10; i++ {
		if i == 0 {
			slice = append(slice, i+10)
		} else {
			slice = append(slice, i*10)
		}

	}

	fmt.Println("1ª questão")
	fmt.Printf("%v\n\n", slice[:3])

	fmt.Println("2ª questão")
	fmt.Printf("%v\n\n", slice[4:])

	fmt.Println("3ª questão")
	fmt.Printf("%v\n\n", slice[1:7])

	// fmt.Println(len(slice))

	fmt.Println("4ª questão")
	fmt.Printf("%v\n\n", slice[2:9])

	fmt.Println("5ª questão")
	fmt.Printf("%v\n\n", slice[2:len(slice)-1])

	fmt.Printf("\nDatatype: %T\n", slice)

	fmt.Println("")
	fmt.Println("")
}
