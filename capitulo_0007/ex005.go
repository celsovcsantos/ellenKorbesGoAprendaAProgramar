package capitulo_0007

import (
	"fmt"
)

func Ex005() {

	fmt.Println("Exercício 005 - OK")

	fmt.Println("1ª questão")
	for i := 10; i <= 100; i++ {
		restoDivisao := i % 4
		fmt.Printf("%v -> %d\n", i, restoDivisao)
	}
	fmt.Println("")
	fmt.Println("")
}
