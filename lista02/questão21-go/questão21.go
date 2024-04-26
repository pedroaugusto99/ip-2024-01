Questão 21
	
package main

import "fmt"


func main() {
  for {
    // Ler o tamanho da próxima sequência
    var n int
    if _, err := fmt.Scan(&n); err != nil {
      // Se erro ou EOF, terminar o loop
      break
    }

    // Se o tamanho for zero, termina
    if n == 0 {
      break
    }

    // Ler a sequência de números
    num := make([]float64, n)
    for i := 0; i < n; i++ {
      _, err := fmt.Scan(&num[i])
      if err != nil {
        break
      }
    }

    // Verificar se está ordenada
    Ordenada := true
    for i := 0; i < len(num); i++ {
      if num[i] >= num[i+1] {
        Ordenada = false
        break
      }
    }

    // Imprimir a resposta conforme instruído
    if Ordenada {
      fmt.Println("ORDENADA")
    } else  {
      fmt.Println("DESORDENADA")
    } 
  }
}
