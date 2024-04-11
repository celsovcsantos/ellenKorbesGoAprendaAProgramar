package capitulo_0007

import (
	"fmt"
	"time"
)

func Ex004() {

	fmt.Println("Exercício 004 - OK")

	cont := 0
	i := 1985
	anoAtual := time.Now().Year()

	fmt.Println("1ª questão")
	for {
		if i <= anoAtual {
			fmt.Printf("%d ", i)
			cont++
			i++
		} else if i > anoAtual {
			fmt.Println("")
			fmt.Printf("Saindo do loop no i = %d\n", i)
			break
		}
	}
	fmt.Println("")

	fmt.Printf("Qtd de anos: %d", cont)

	fmt.Println("")
	fmt.Println("")
}
