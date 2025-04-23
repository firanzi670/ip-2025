package main

import (
	"fmt"
)

func main() {
	var vetor1 [10]int
	var vetor2 [5]int

	
	fmt.Println("Digite 10 números inteiros para o primeiro vetor:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scan(&vetor1[i])
	}

	
	fmt.Println("\nDigite 5 números inteiros para o segundo vetor:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scan(&vetor2[i])
	}

	
	somaVetor2 := 0
	for _, valor := range vetor2 {
		somaVetor2 += valor
	}

	
	var paresResultado []int
	var imparesResultado []int

	for _, valor := range vetor1 {
		if valor%2 == 0 {
			paresResultado = append(paresResultado, valor+somaVetor2)
		} else {
			imparesResultado = append(imparesResultado, valor+somaVetor2)
		}
	}

	
	fmt.Println("\nPrimeiro vetor resultante (pares + soma segundo vetor):", paresResultado)
	fmt.Println("Segundo vetor resultante (ímpares + soma segundo vetor):", imparesResultado)
}
