package capitulo_0003

import "fmt"

// 002 - go run capitulo_0003/ex002_main.go

//package level scope
var xEx002 int
var yEx002 string
var zEx002 bool

func Ex002() {

	fmt.Println("Exercício 002 - OK")

	//1ª questão
	fmt.Println("1ª questão")
	fmt.Printf("x -> %v -> %T\n", xEx002, xEx002)
	fmt.Printf("y -> '%v' -> %T\n", yEx002, yEx002)
	fmt.Printf("z -> %v -> %T\n", zEx002, zEx002)
	//2ª questão -
	fmt.Println("2ª questão")
	fmt.Println("Resposta: valores zero")

	fmt.Println("")
	fmt.Println("")
}
