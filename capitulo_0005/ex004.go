package capitulo_0005

import "fmt"

func Ex004() {

	x := 200

	fmt.Println("Exercício 004 - OK")
	fmt.Print("1ª questão:\nAtribuição de variável\n")
	fmt.Println("2ª questão: ")
	fmt.Printf("x -> dec: %d\tbin: %b\thexa: %#x\n", x, x, x)
	fmt.Print("3ª questão:\nDeslocamento de bits\n")
	y := x << 1
	fmt.Println("4ª questão:")
	fmt.Printf("y -> dec: %d\tbin: %b\thexa: %#x\n", y, y, y)

	fmt.Printf("")
	fmt.Printf("")
}
