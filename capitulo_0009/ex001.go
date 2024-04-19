package capitulo_0009

import "fmt"

// 001 - go run capitulo_0003/ex001_main.go
func Ex001() {
	fmt.Println("Exercício 001 - OK")

	array := [5]int{}

	for i := 0; i < len(array); i++ {
		array[i] = i + 10
	}
	//1ª questão
	fmt.Println("1ª questão")
	for i, value := range array {
		fmt.Println(i, value)
	}

	fmt.Printf("Datatype: %T", array)

	fmt.Println("")
	fmt.Println("")
}
