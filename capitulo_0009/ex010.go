package capitulo_0009

import "fmt"

// import (
// 	"fmt"
// )

func Ex010() {
	fmt.Println("Exercício 010 - OK")
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

	mapa["teste4"] = []string{"foo", "bar", "baz"}

	for k, value := range mapa {
		fmt.Printf("%v\n", k)
		for i, v := range value {
			fmt.Printf("\t%v - %v\n", i, v)
		}
	}

	fmt.Println("")
	fmt.Println("")

	delete(mapa, "teste2")

	for k, value := range mapa {
		fmt.Printf("%v\n", k)
		for i, v := range value {
			fmt.Printf("\t%v - %v\n", i, v)
		}
	}

	fmt.Println("")
	fmt.Println("")
}
