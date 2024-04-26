Questão 12

package main

import ("fmt"
"math")

func calcularIngressoEsperado(valorIngresso, valorAtual float64) int {
	delta := (valorIngresso - valorAtual) / 0.50
	
	if delta > 0 {
		return 120 + 25*int(delta)
	} else {
		return 120 + 30*int(delta)
	}
}

func calcularLucro(valorIngresso, valorAtual float64, ingressos int) float64 {
	despesasFixas := 200.0
	despesasVariaveis := 0.05 * float64(ingressos)
	receita := valorAtual * float64(ingressos)
	return receita - (despesasFixas + despesasVariaveis)
}

func main() {
	var valorIngresso, valorInicial, valorFinal float64

	fmt.Scan(&valorIngresso, &valorInicial, &valorFinal)

	if valorInicial >= valorFinal {
		fmt.Println("INTERVALO INVALIDO.")
		return
	}

	melhorLucro := -math.MaxFloat64
	var melhorValor float64
	var melhorIngressos int

	for valor := valorInicial; valor <= valorFinal; valor++ {
		ingressos := calcularIngressoEsperado(valorIngresso, valor)
		lucro := calcularLucro(valorIngresso, valor, ingressos)

		fmt.Printf("V: %.2f, N: %d, L: %.2f\n", valor, ingressos, lucro)

		if lucro > melhorLucro {
			melhorLucro = lucro
			melhorValor = valor
			melhorIngressos = ingressos
		}
	}

	if melhorLucro < 0 {
		melhorLucro = 0
		melhorValor = 0
		melhorIngressos = 0
	}

	fmt.Printf("Melhor valor final: %.2f\n", melhorValor)
	fmt.Printf("Lucro: %.2f\n", melhorLucro)
	fmt.Printf("Numero de ingressos: %d\n", melhorIngressos)
}
