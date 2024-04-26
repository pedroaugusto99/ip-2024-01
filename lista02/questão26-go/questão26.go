Questão 26

package main

import (
	"fmt"
	"math"
)

// Função para calcular o fatorial de um número
func fatorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * fatorial(n-1)
}

// Função para calcular o seno usando a série de Taylor
func senoTaylor(x float64, n int) float64 {
	seno := 0.0
	for i := 0; i <= n; i++ {
		termo := math.Pow(-1, float64(i)) * math.Pow(x, float64(2*i+1)) / float64(fatorial(2*i+1))
		seno += termo
	}
	return seno
}

func main() {
	var x float64
	var n int
	fmt.Scan(&x, &n)

	// Calcular o seno de x usando a série de Taylor com N termos
	sin := senoTaylor(x, n)

	// Imprimir o resultado no formato especificado
	fmt.Printf("seno(%.2f) = %.6f\n", x, sin)
}
