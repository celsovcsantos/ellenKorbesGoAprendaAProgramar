package capitulo_0009

import "fmt"

func Ex002() {

	fmt.Println("Exercício 002 - OK")

	slice := []int{}

	for i := 0; i < 10; i++ {
		slice = append(slice, i+10)
	}

	//1ª questão
	fmt.Println("1ª questão")
	for i, v := range slice {
		fmt.Printf("%d\t%d\n", i, v)
	}

	fmt.Printf("Datatype: %T\n", slice)

	fmt.Println("")
	fmt.Println("")
}
