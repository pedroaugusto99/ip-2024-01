package main

import "fmt"

type VetorInfo struct {
	MaiorElemento int
	IndiceMaior   int
}

func main() {
	var N int
	var vetores []VetorInfo
	
	fmt.Println("Entrada")
	for {
		fmt.Scan(&N)
		if N == 0 {
			break
		}

		vetor := make([]int, N)
		for i := range vetor {
			fmt.Scan(&vetor[i])
		}

		maiorElemento := vetor[0]
		indiceMaior := 0

		for i, elemento := range vetor {
			if elemento > maiorElemento {
				maiorElemento = elemento
				indiceMaior = i
			}
		}

		vetores = append(vetores, VetorInfo{
			MaiorElemento: maiorElemento,
			IndiceMaior:   indiceMaior,
		})
	}
    fmt.Println("Saída")
	for _, info := range vetores {
		fmt.Printf("%d %d\n", info.IndiceMaior, info.MaiorElemento) // +1 porque os índices começam em 0
	}
}
