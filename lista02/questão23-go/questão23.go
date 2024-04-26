Questão 23

package main

import (
	"fmt"
)

// Função para encontrar todos os divisores de um número
func encontrarDivisores(n int) []int {
	var divisores []int
	for i := 1; i <= n/2; i++ { // Divisores são todos os números até n/2
		if n%i == 0 {
			divisores = append(divisores, i)
		}
	}
	return divisores
}

// Função para calcular a soma dos divisores
func somaDivisores(divisores []int) int {
	soma := 0
	for _, divisor := range divisores {
		soma += divisor
	}
	return soma
}

func main() {
	var n int
	fmt.Scan(&n)

	encontrados := 0
	paresAmigos := make(map[int]bool) // Mapeamento para evitar duplicidades
	num := 1

	for encontrados < n {
		divisores := encontrarDivisores(num)
		soma := somaDivisores(divisores)

		// Verificar se são amigos e se não são repetidos
		if soma > num { // Garante que o menor número vem primeiro
			divisoresSoma := encontrarDivisores(soma)
			somaDoOutro := somaDivisores(divisoresSoma)
			if somaDoOutro == num && !paresAmigos[num] && !paresAmigos[soma] {
				fmt.Printf("(%d,%d)\n", num, soma)
				paresAmigos[num] = true
				paresAmigos[soma] = true
				encontrados++
			}
		}

		num++ // Incrementa o número para buscar o próximo par de números amigos
	}
}
