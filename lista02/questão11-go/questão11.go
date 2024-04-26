Questão 11

package main

import (
  "fmt"
  "math"
)

func main() {
  var n, e float64
  fmt.Print("Insira o valor de n: ")
  fmt.Scan(&n)
  fmt.Print("Insira o valor do erro tolerável e: ")
  fmt.Scan(&e)

  
  r := 1.0
  i := 0

  
  for {
    i++
    
    r = (r + n/r) / 2
     
    erro := math.Abs(n - r*r)

    
    fmt.Printf("r = %.9f, Erro = %.9f\n", r, erro)

  
    if erro <= e {
      break
    }
  }
}
