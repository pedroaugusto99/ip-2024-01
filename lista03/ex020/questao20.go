package main

import (
  "fmt"
  "math"
)

type Ponto struct {
  X, Y, Z float64
}

func calcularDistancia(ponto1, ponto2 Ponto) float64 {
  diferencaX := math.Abs(ponto2.X - ponto1.X)
  diferencaY := math.Abs(ponto2.Y - ponto1.Y)
  diferencaZ := math.Abs(ponto2.Z - ponto1.Z)

  if diferencaX > diferencaY && diferencaX > diferencaZ {
    return diferencaX
  } else if diferencaY > diferencaX && diferencaY > diferencaZ {
    return diferencaY
  } else {
    return diferencaZ
  }
}


func main() {
  var N int
  
  fmt.Scanln(&N)
  
  numeros := make([]Ponto, N)
  for i := 0; i < N; i++ {
    fmt.Scan(&numeros[i].X, &numeros[i].Y, &numeros[i].Z)
  }
  
  for i := 0; i < N-1; i++ {
    distancia := calcularDistancia(numeros[i], numeros[i+1])
    fmt.Printf("%.2f\n", distancia)
}
}
