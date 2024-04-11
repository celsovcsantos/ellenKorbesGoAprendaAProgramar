package capitulo_0007

import "fmt"

func Ex002() {

	fmt.Println("Exercício 002 - OK")

	//1ª questão
	fmt.Println("1ª questão")
	for i := 65; i <= 90; i++ {
		fmt.Printf("\t%v\n", i)
		for x := 1; x <= 3; x++ {
			fmt.Printf("\t\t%#U\n", i)
		}
	}

	fmt.Println("")
	fmt.Println("")
}
