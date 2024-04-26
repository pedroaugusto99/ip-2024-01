Questão 27

package main

import (
	"fmt"
)

// Função para verificar se um número é primo
func numPrimo(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Função para fatorar um número em fatores primos
func primoFatores(n int) []int {
	fatores := []int{}
	if n <= 1 {
		return fatores
	}

	// Divisão por 2 até que não seja mais divisível
	for n%2 == 0 {
		fatores = append(fatores, 2)
		n /= 2
	}

	// Testar todos os números ímpares de 3 até √n
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			fatores = append(fatores, i)
			n /= i
		}
	}

	// Se sobrou algum número maior que 1, é um fator primo
	if n > 1 {
		fatores = append(fatores, n)
	}

	return fatores
}

func main() {
	var n int
	fmt.Scan(&n)

	if n <= 1 {
		fmt.Printf("Fatoracao nao e possivel para o numero %d!\n", n)
		return
	}

	fatore := primoFatores(n)

	if len(fatores) == 0 {
		fmt.Printf("Fatoracao nao e possivel para o numero %d!\n", n)
	} else {
		fmt.Printf("%d = ", n)
		for i, fatores := range fatores {
			if i > 0 {
				fmt.Print(" x ")
			}
			fmt.Print(factor)
		}
		fmt.Println() // para a quebra de linha ao final
	}
}
