package capitulo_0009

import "fmt"

// import (
// 	"fmt"
// )

func Ex007() {
	fmt.Println("Exercício 007 - OK")

	fmt.Println("1ª questão")
	slice := [][]string{{"Celso", "Santos", "videogame"}, {"Michelle", "Alves", "Viajar"}, {"Matheus", "Santos", "futebol"}}

	for _, v := range slice {
		fmt.Println(v)
	}

	fmt.Println("")
	fmt.Println("")

	for _, v := range slice {
		fmt.Println(v[0])
		for i, itemV := range v {
			if i > 0 {
				fmt.Printf("\t%v\n", itemV)
			}
		}
	}

	fmt.Println("")
	fmt.Println("")
}
