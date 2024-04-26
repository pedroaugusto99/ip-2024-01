Questão 24

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

// Função para calcular o cosseno usando a série de Taylor
func cosTaylor(x float64, n int) float64 {
	cos := 0.0
	for i := 0; i <= n; i++ { // N é o número de termos
		termo := math.Pow(-1, float64(i)) * math.Pow(x, float64(2*i)) / float64(fatorial(2*i))
		cos += termo
	}
	return cos
}

func main() {
	var x float64
	var n int
	fmt.Scan(&x, &n)

	// Calcular o cosseno de x usando a série de Taylor com N termos
	cos := cosTaylor(x, n)

	// Imprimir o resultado no formato especificado
	fmt.Printf("cos(%.2f) = %.6f\n", x, cos)
}
