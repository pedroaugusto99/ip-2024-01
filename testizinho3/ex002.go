package main

import "fmt"

func troca(x []int, i, j int) {
	x[i], x[j] = x[j], x[i]
}

func trocaOpostosSeMenor(x []int) {
	tam := len(x)
	for i := 0; i < tam/2; i++ {
		contrário := tam - 1 - i
		if x[contrário] < x[i] {
			troca(x, i, contrário)
		}
	}
}

func main() {
	var testes, N int
	fmt.Scanln(&testes)
	for t := 0; t < testes; t++ {
		fmt.Scanln(&N)
		if N > 500 {
			continue
		}
		x := make([]int, N)
		for i := range x {
			fmt.Scan(&x[i])
		}
		trocaOpostosSeMenor(x)
		fmt.Println(x)
	}
}
