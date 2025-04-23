package main

import (
	"fmt"
)

func main() {
	var numeros [10]int
	var pares []int
	var impares []int
	somaPares := 0

	
	fmt.Println("Digite 10 números inteiros:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scan(&numeros[i])
	}

	
	for _, num := range numeros {
		if num%2 == 0 {
			pares = append(pares, num)
			somaPares += num
		} else {
			impares = append(impares, num)
		}
	}

	
	fmt.Println("\n=== RESULTADOS ===")
	fmt.Println("a) Números pares digitados:", pares)
	fmt.Println("b) Soma dos números pares:", somaPares)
	fmt.Println("c) Números ímpares digitados:", impares)
	fmt.Println("d) Quantidade de números ímpares:", len(impares))
}
