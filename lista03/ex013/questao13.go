package main

import "fmt"


func main() {
	
	var N int
	fmt.Println("Entrada")
	fmt.Scan(&N)

	frequencias := make(map[int]int)

	for i := 0; i < N; i++ {
		var numero int
		fmt.Scan(&numero)
		frequencias[numero]++
	}

	maiorFrequencia := 0
	numeroComMaiorFrequencia := 0
	for numero, frequencia := range frequencias {
		if frequencia > maiorFrequencia {
			maiorFrequencia = frequencia
			numeroComMaiorFrequencia = numero
		}
	}

	var menorNumeroComMaiorFrequencia int
	for numero, frequencia := range frequencias {
		if frequencia == maiorFrequencia && numero < numeroComMaiorFrequencia {
			menorNumeroComMaiorFrequencia = numero
			fmt.Println("Menor Número:",menorNumeroComMaiorFrequencia)
		}
	}

    fmt.Println("")
    fmt.Println("Saída")
	fmt.Println(numeroComMaiorFrequencia)
	fmt.Println(maiorFrequencia)
}
