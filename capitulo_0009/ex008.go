package capitulo_0009

import "fmt"

// import (
// 	"fmt"
// )

func Ex008() {
	fmt.Println("Exercício 008 - OK")

	fmt.Println("1ª questão")

	mapa := map[string][]string{
		"teste1": {"fut", "vol", "nat"},
		"teste2": {"bic", "pet", "cor"},
		"teste3": {"fil", "vid", "cam"},
	}

	for k, value := range mapa {
		fmt.Printf("%v\n", k)
		for i, v := range value {
			fmt.Printf("\t%v - %v\n", i, v)
		}
	}

	fmt.Println("")
	fmt.Println("")
}
