package main

import (
	"fmt"
)

func main() {
	var numeros [10]int
	var encontrou bool = false

	
	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o %dº número inteiro: ", i+1)
		fmt.Scan(&numeros[i])
	}

	
	fmt.Println("\nNúmeros superiores a 50 e suas posições:")
	for i, num := range numeros {
		if num > 50 {
			fmt.Printf("Valor: %d | Posição: %d\n", num, i)
			encontrou = true
		}
	}

	
	if !encontrou {
		fmt.Println("Nenhum número superior a 50 foi encontrado.")
	}
}
