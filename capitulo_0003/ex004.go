package capitulo_0003

import "fmt"

type tipoEx004 int

var xEx004 tipoEx004

func Ex004() {
	fmt.Println("Exercício 004 - OK")
	fmt.Print("1ª questão: ")
	fmt.Printf("x -> %v\n", xEx004)
	fmt.Print("2ª questão: ")
	fmt.Printf("x -> %T\n", xEx004)
	fmt.Printf("3ª questão: atribuindo o valor 42\n")
	xEx004 = 42
	fmt.Print("4ª questão: ")
	fmt.Printf("x -> %v\n", xEx004)
}
