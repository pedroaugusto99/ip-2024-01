Questão 25

package main

import (
	"fmt"
	"math"
)

// Função para calcular o fatorial de um número
func fatorial(n int) int {
	resultado := 1
	for i := 2; i <= n; i++ {
		resultado *= i
	}
	return resultado
}

// Função para calcular o valor de e^x usando a série de Taylor até N termos
func euler(x float64, n int) float64 {
	euler := 0.0
	for i := 0; i <= n; i++ {
		termo := math.Pow(x, float64(i)) / float64(fatorial(i))
		euler += termo
	}
	return euler
}

func main() {
	var x float64
	var n int
	fmt.Scan(&x, &n)

	// Calcular e^x usando a série de Taylor até N termos
	eulerVal := euler(x, n)

	// Imprimir o resultado com formato especificado
	fmt.Printf("e^%.2f = %.6f\n", x, eulerVal)
}
