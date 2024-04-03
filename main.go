package main

import "fmt"

// Aula2.2
// func main() {
// 	numeroDeBytes, erro := fmt.Println("Hello, World", "teste", 100)
// 	fmt.Println(numeroDeBytes, erro)
// }

// Aula2.3
var y = "Olá, bom dia!"

func main() {
	x := 10.0

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	x, z := 20, 40
	fmt.Printf("x: %v, %T\n", x, x)
	teste, erro := fmt.Printf("z: %v, %T\n", z, z)
	fmt.Println(teste, erro)
}
