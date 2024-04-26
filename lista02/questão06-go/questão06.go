Questão 06

package main

import (
	"fmt"
)

// Função para encontrar a maior subsequência de elementos iguais consecutivos
func maxConsecutiveEqualElements(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	maxLength := 1
	currentLength := 1

	// Percorre a sequência de números
	for i := 1; i < len(numbers); i++ {
		if numbers[i] == numbers[i-1] {
			// Se o número atual é igual ao anterior, aumente a contagem atual
			currentLength++
		} else {
			// Se a sequência é interrompida, compare com a máxima encontrada
			if currentLength > maxLength {
				maxLength = currentLength
			}
			currentLength = 1 // reiniciar a contagem
		}
	}

	// Verifique se a última subsequência é a mais longa
	if currentLength > maxLength {
		maxLength = currentLength
	}

	return maxLength
}

func main() {
	var n int
	fmt.Scan(&n) // Ler o número de elementos na sequência

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i]) // Ler a sequência de números
	}

	maxSubsequence := maxConsecutiveEqualElements(numbers)

	fmt.Printf("A maior subsequência de elementos iguais possui %d elementos.\n", maxSubsequence)
}
