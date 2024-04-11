package capitulo_0007

import "fmt"

// import (
// 	"fmt"
// )

func Ex006() {
	fmt.Println("Exercício 006 - OK")

	fmt.Println("1ª questão")
	for i := 10; i <= 100; i++ {
		if i == 50 {
			restoDivisao := i % 4
			fmt.Printf("%v -> %d\n", i, restoDivisao)
		}
	}
	fmt.Println("")
	fmt.Println("")
}
