package capitulo_0009

import (
	"fmt"
)

// import (
// 	"fmt"
// )

func Ex006() {
	fmt.Println("Exercício 006 - OK")

	fmt.Println("1ª questão: slice := make([]string, 26, 26)")
	//slice := make([]string, 0, 26)
	//slice = append(slice, "Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins")
	slice := []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	fmt.Println("")

	// fmt.Printf("%v\n", slice)
	fmt.Println("2ª questão")
	fmt.Printf("Len: %d\n", len(slice))
	fmt.Printf("Cap: %d\n", cap(slice))

	fmt.Println("")

	fmt.Println("3ª questão")
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	fmt.Println("")
	fmt.Println("")
}
