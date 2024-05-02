package main

import (
	"fmt"
	"math"
)

type Ponto struct {
	X, Y, Z float64
}

func calcularDistancia(ponto1, ponto2 Ponto) float64 {
	diferencaX := ponto2.X - ponto1.X
	diferencaY := ponto2.Y - ponto1.Y
	diferencaZ := ponto2.Z - ponto1.Z

	return math.Sqrt(diferencaX*diferencaX + diferencaY*diferencaY + diferencaZ*diferencaZ)
}

func main() {
	var N int
	fmt.Scan(&N)

	pontos := make([]Ponto, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&pontos[i].X, &pontos[i].Y, &pontos[i].Z)
	}

	for i := 0; i < N-1; i++ {
		distancia := calcularDistancia(pontos[i], pontos[i+1])
		fmt.Printf("%.2f\n", distancia)
	}
}

