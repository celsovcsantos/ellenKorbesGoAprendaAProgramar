package capitulo_0007

import "fmt"

// import (
// 	"fmt"
// )

func Ex007() {
	fmt.Println("Exercício 007 - OK")

	fmt.Println("1ª questão")
	for i := 10; i <= 100; i++ {
		restoDivisao := i % 4
		if i == 50 {
			fmt.Printf("%v -> %d\n", i, restoDivisao)
		} else if i == 74 {
			fmt.Printf("%v -> %d\n", i, restoDivisao)
		} else {
			fmt.Printf("%v -> %d\n", i, 10)
		}
	}
	fmt.Println("")
	fmt.Println("")
}
