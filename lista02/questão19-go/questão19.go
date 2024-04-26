Questão 19

package main

import "fmt"

func main() {
  var m int
  
  fmt.Print("Digite um número inteiro maior que zero: ")
  fmt.Scanln(&m)
  
  if m <= 0 {
    fmt.Println("Dígito Inválido")
  }
  
  for k := 1; k <= m; k++ {
    soma := 0
    inicio := k*k - k + 1 // O primeiro ímpar da sequência
    sequencia := make([]int, k) // Para armazenar a sequência de ímpares

    for i := 0; i < k; i++ {
      sequencia[i] = inicio + 2*i // Gera os números ímpares consecutivos
      soma += sequencia[i] // Calcula a soma
    }

    // Formatar a saída
    fmt.Printf("%d×%d×%d = ", k, k, k)
    for i, v := range sequencia {
      if i == len(sequencia)-1 {
        fmt.Printf("%d\n", v) // Sem + no final
      } else {
        fmt.Printf("%d+", v) // Com + para intermediários
      }
    }
}
}
