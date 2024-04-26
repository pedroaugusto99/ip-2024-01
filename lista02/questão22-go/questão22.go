Questão 22

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Função para calcular o Máximo Divisor Comum (MDC) usando o algoritmo de Euclides
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// Função para converter um número decimal para uma fração simplificada
func decimalParaFracao(n float64) (int, int) {
	// Convertendo o número para string para analisar a parte decimal
	nStr := fmt.Sprintf("%.10f", n) // Limitar a 10 casas decimais
	nStr = strings.TrimRight(nStr, "0") // Remover zeros à direita
	parts := strings.Split(nStr, ".")

	// Parte inteira do número
	integral, _ := strconv.Atoi(parts[0])

	// Parte decimal
	if len(parts) < 2 {
		// Não há parte decimal, então é uma fração com denominador 1
		return integral, 1
	}

	decimalPart := parts[1]
	decimalDigits := len(decimalPart)

	// Denominador é 10^decimalDigits
	denominator := int(math.Pow(10, float64(decimalDigits)))
	// Numerador é a parte inteira vezes o denominador mais a parte decimal
	decimalInt, err := strconv.Atoi(decimalPart) // Trata o erro ao converter a parte decimal para número
	if err != nil {
		fmt.Println("Erro ao converter parte decimal para número:", err)
		return 0, 0
	}
	numerator := integral*denominator + decimalInt

	// Obter o MDC para simplificar a fração
	mdc := gcd(numerator, denominator)

	// Retornar a fração simplificada
	return numerator / mdc, denominator / mdc
}

func main() {
	var n float64
	fmt.Scan(&n)

	num, den := decimalParaFracao(n)

	fmt.Printf("%d/%d\n", num, den)
}
