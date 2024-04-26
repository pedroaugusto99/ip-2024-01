Questão 14

package main

import "fmt"

func main() {
  var m, n int

  fmt.Print("Digite o número de linhas (m): ")
  fmt.Scan(&m)
  fmt.Print("Digite o número de colunas (n): ")
  fmt.Scan(&n)

  if m <= 0 || n <= 0 {
    fmt.Println("Dimensões inválidas. Os valores de m e n devem ser maiores que zero.")
    return
  }

  for i := 1; i <= m; i++ {
    
    for j := 1; j <= n; j++ {
       
      if i > j {
        fmt.Printf("(%d-%d)\n—", i, j)
      }
    }
  }
}
