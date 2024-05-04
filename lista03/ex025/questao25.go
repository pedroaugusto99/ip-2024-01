package main

import "fmt"

func main() {
  var sorteados [6]int
  for i := 0; i < 6; i++ {
    fmt.Scan(&sorteados[i])
  }

  var sorteios int
  fmt.Scan(&sorteios)

  sena := 0
  quina := 0
  quadra := 0
  
  for i := 0; i < sorteios; i++ {
    var bilhetes [6]int
    for j := 0; j < 6; j++ {
      fmt.Scan(&bilhetes[j])
    }

    acertos := 0
    for _, dezena := range bilhetes {
      for _, sorteado := range sorteados {
        if dezena == sorteado {
          acertos++
          break
        }
      }
    }

    switch acertos {
    case 6:
      sena++
    case 5:
      quina++
    case 4:
      quadra++
    }
  }

  if sena > 0 {
    fmt.Printf("Houve %d acertador(es) da sena\n", sena)
  } else {
    fmt.Println("Nao houve acertador para sena")
  }

  if quina > 0 {
    fmt.Printf("Houve %d acertador(es) da quina\n", quina)
  } else {
    fmt.Println("Nao houve acertador para quina")
  }

  if quadra > 0 {
    fmt.Printf("Houve %d acertador(es) da quadra\n", quadra)
  } else {
    fmt.Println("Nao houve acertador para quadra")
  }
}
