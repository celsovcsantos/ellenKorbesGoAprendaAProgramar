package capitulo_0005

import "fmt"

func Ex002() {

	var1 := 10
	var2 := 10
	var3 := 20

	fmt.Println("Exercício 002 - OK")

	res1 := (var1 == var2)
	res2 := (var1 != var3)
	res3 := (var3 <= var2)
	res4 := (var2 < var3)
	res5 := (var3 >= var2)
	res6 := (var2 > var1)

	//1ª questão
	fmt.Println("1ª questão")
	fmt.Printf("%v == %v -> %v\n", var1, var2, res1)
	fmt.Printf("%v != %v -> %v\n", var1, var3, res2)
	fmt.Printf("%v <= %v -> %v\n", var3, var2, res3)
	fmt.Printf("%v < %v -> %v\n", var2, var3, res4)
	fmt.Printf("%v >= %v -> %v\n", var3, var2, res5)
	fmt.Printf("%v > %v -> %v\n", var2, var1, res6)

	fmt.Println("")
	fmt.Println("")
}
