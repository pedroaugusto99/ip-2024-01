Questão 15

package main

import "fmt"

// Função para reverter a ordem dos caracteres de uma string
func reverse(s string) string {
  n := len(s)
  rev := make([]byte, n)
  for i := 0; i < n; i++ {
    rev[n-1-i] = s[i]
  }
  return string(rev)
}

func main() {
  var t int
  fmt.Scan(&t) // Lê o número de casos de teste

  for i := 0; i < t; i++ {
    var A, B string
    fmt.Scan(&A, &B) // Lê dois números de três dígitos como strings

    // Inverte as strings para obter o número "lido ao contrário"
    Ainv := reverse(A)
    Binv := reverse(B)

    // Compara as strings invertidas
    if Ainv > Binv {
      fmt.Println(Ainv) // Se A invertido for maior, imprime A invertido
    } else {
      fmt.Println(Binv) // Se B invertido for maior, imprime B invertido
    }
  }
}
