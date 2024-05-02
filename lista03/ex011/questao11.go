package main

import "fmt"

func main() {
  var n int
  fmt.Println("Entrada")
  fmt.Scan(&n)

  numeros := make([]int,n,n)
  for i := range numeros {
    fmt.Scan(&numeros[i])
  }
  
  maiorValor := numeros[0]
    menorValor := numeros[0]
    for _, numero := range numeros {
        if numero > maiorValor {
            maiorValor = numero
        }
        if numero < menorValor {
            menorValor = numero
        }
    }
    
  fmt.Println("Saída")
  for i := n - 1; i >= 0; i-- {
    fmt.Printf("%v ",numeros[i])
  }
  fmt.Println("")
  fmt.Println(maiorValor)
  fmt.Println(menorValor)
}
