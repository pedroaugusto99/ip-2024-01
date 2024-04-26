Questão 18

package main

import (
  "fmt"
  "math"
)

// Função para verificar se um número é primo
func numprimo(n int) bool {
  if n <= 1 {
    return false
  }
  for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
    if n%i == 0 {
      return false
    }
  }
  return true
}

// Função para obter o próximo número primo após um dado número
func proximoPrimo(n int) int {
  for i := n + 1; ; i++ {
    if numprimo(i) {
      return i
    }
  }
}

// Função para calcular o MMC usando fatoração
func calcularMMC(a, b, c int) int {
  n1, n2, n3 := a, b, c
  mmc := 1
  primo := 2

  fmt.Printf("%2d %2d %2d :", n1, n2, n3)

  // Enquanto pelo menos um dos números não for 1
  for n1 != 1 || n2 != 1 || n3 != 1 {
    dividiu := false

    if n1%primo == 0 {
      n1 /= primo
      dividiu = true
    }

    if n2%primo == 0 {
      n2 /= primo
      dividiu = true
    }

    if n3%primo == 0 {
      n3 /= primo
      dividiu = true
    }

    if dividiu {
      fmt.Printf(" %2d\n", primo) // Exibe o primo usado para fatorar
      fmt.Printf("%2d %2d %2d :", n1, n2, n3)
      mmc *= primo
    } else {
      primo = proximoPrimo(primo) // Obter o próximo primo
    }
  }

  return mmc
}

func main() {
  var a, b, c int

  // Leitura dos três números
  fmt.Println("Digite três números inteiros:")
  fmt.Scan(&a, &b, &c)

  // Calcular o MMC usando a fatoração
  fmt.Println("Processo de fatoração para calcular o MMC:")
  mmc := calcularMMC(a, b, c)

  fmt.Printf("\nMMC de %d, %d e %d é: %d\n", a, b, c, mmc)
}

