package main

import (
	"fmt"
)

func lerConjunto(tamanho int) map[int]bool {
	conjunto := make(map[int]bool)
	for i := 0; i < tamanho; i++ {
		var num int
		fmt.Scan(&num)
		conjunto[num] = true
	}
	return conjunto
}

func uniao(conjuntoA, conjuntoB map[int]bool) map[int]bool {
	conjuntoUniao := make(map[int]bool)
	for chave := range conjuntoA {
		conjuntoUniao[chave] = true
	}
	for chave := range conjuntoB {
		conjuntoUniao[chave] = true
	}
	return conjuntoUniao
}

func interseccao(conjuntoA, conjuntoB map[int]bool) map[int]bool {
	conjuntoInterseccao := make(map[int]bool)
	for chave := range conjuntoA {
		if _, existe := conjuntoB[chave]; existe {
			conjuntoInterseccao[chave] = true
		}
	}
	return conjuntoInterseccao
}

func imprimirConjunto(conjunto map[int]bool) {
	fmt.Print("(")
	primeiro := true
	for chave := range conjunto {
		if !primeiro {
			fmt.Print(",")
		}
		fmt.Print(chave)
		primeiro = false
	}
	fmt.Println(")")
}

func main() {
	var TA, TB int
	fmt.Scan(&TA)
	fmt.Scan(&TB)

	conjuntoA := lerConjunto(TA)
	conjuntoB := lerConjunto(TB)

	conjuntoUniao := uniao(conjuntoA, conjuntoB)
	conjuntoInterseccao := interseccao(conjuntoA, conjuntoB)

	imprimirConjunto(conjuntoUniao)
	imprimirConjunto(conjuntoInterseccao)
}
