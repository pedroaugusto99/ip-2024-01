Questão 16

package main

import (
	"fmt"
	"sort"
)

type Cateto struct {
	a, b int
}

// Encontrar pares de catetos para uma dada hipotenusa
func encontrarCatetos(hip int) []Cateto {
	hip2 := hip * hip
	catetos := []Cateto{}
	for a := 1; a < hip; a++ { // a deve ser menor que a hipotenusa
		for b := a; b < hip; b++ { // b deve ser maior ou igual a a, mas menor que a hipotenusa
			if a*a+b*b == hip2 { // Teste do Teorema de Pitágoras
				catetos = append(catetos, Cateto{a, b})
			}
		}
	}
	return catetos
}

func main() {
	var n int
	fmt.Scan(&n) // Ler o valor de n

	for h := 1; h <= n; h++ { // Iterar sobre as hipotenusas de 1 a n
		catetos := encontrarCatetos(h) // Encontrar catetos para uma hipotenusa
		if len(catetos) > 0 { // Se existirem catetos válidos para esta hipotenusa
			sort.Slice(catetos, func(i, j int) bool { // Ordenar por a, depois por b
				return catetos[i].a < catetos[j].a
			})
			for _, c := range catetos { // Imprimir todos os pares de catetos
				fmt.Printf("hipotenusa = %d, catetos %d e %d\n", h, c.a, c.b)
			}
		}
	}
}
