package capitulo_0007

import "fmt"

// import (
// 	"fmt"
// )

func Ex008() {
	fmt.Println("Exercício 008 - OK")

	fmt.Println("1ª questão")
	for i := 10; i <= 100; i++ {
		restoDivisao := i % 4
		switch {
		case i == 50:
			fmt.Printf("%v -> %d\n", i, restoDivisao)
		case i == 74:
			fmt.Printf("%v -> %d\n", i, restoDivisao)
		default:
			fmt.Printf("%v -> %v\n", i, "def")
		}
	}
	fmt.Println("")
	fmt.Println("")
}
