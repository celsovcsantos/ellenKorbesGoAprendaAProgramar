package capitulo_0007

import (
	"fmt"
	"time"
)

func Ex003() {

	fmt.Println("Exercício 003 - OK")

	cont := 0
	i := 1985
	anoAtual := time.Now().Year()

	fmt.Println("1ª questão")
	for i <= anoAtual {
		fmt.Printf("%d ", i)
		cont++
		i++
	}
	fmt.Println("")

	fmt.Printf("Qtd de anos: %d", cont)

	fmt.Println("")
	fmt.Println("")
}
