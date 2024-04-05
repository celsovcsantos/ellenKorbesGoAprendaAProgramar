package capitulo_0003

import "fmt"

type tipoEx005 int

var xEx005 tipoEx005
var yEx005 int

func Ex005() {
	fmt.Println("Exercício 004 - OK")
	fmt.Print("1ª questão: ")
	fmt.Printf("x -> %v\n", xEx005)
	fmt.Print("2ª questão: ")
	fmt.Printf("x -> %T\n", xEx005)
	fmt.Printf("3ª questão: atribuindo o valor 42\n")
	xEx005 = 42
	fmt.Print("4ª questão: ")
	fmt.Printf("x -> %v\n", xEx005)

	fmt.Println("Até aqui x")
	fmt.Println("A partir daqui y")

	fmt.Println("Exercício 005 - OK")
	fmt.Print("1ª questão: ")
	yEx005 = int(xEx005)
	fmt.Printf("x -> %v / x -> %T\n", xEx005, xEx005)
	fmt.Print("2ª questão: ")
	fmt.Printf("y -> %v\n", yEx005)
	fmt.Print("3ª questão: ")
	fmt.Printf("y -> %T\n", yEx005)
}
