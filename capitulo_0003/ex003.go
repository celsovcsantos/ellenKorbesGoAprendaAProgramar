package capitulo_0003

import "fmt"

// package level scope
var xEx003 = 42
var yEx003 = "James Bond"
var zEx003 = true

func Ex003() {

	fmt.Println("Exercício 003 - OK")

	fmt.Printf("1ª questão\natribuição de variáveis\n")
	s := fmt.Sprintf("x = %v / y = %v / z = %v", xEx003, yEx003, zEx003)
	fmt.Println("2ª questão")
	fmt.Println(s)

	fmt.Println("")
	fmt.Println("")
}
