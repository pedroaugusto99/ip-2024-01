package main

import "fmt"


func main() {
  var cpf int
  fmt.Scan(&cpf)

  resultado := make([]string, cpf)

  for i := 0; i < cpf; i++ {
    x := make([]int, 11)
    for j := range x {
      fmt.Scan(&x[j])
    }

    b1 := (x[0]*1 + x[1]*2 + x[2]*3 + x[3]*4 + x[4]*5 + x[5]*6 + x[6]*7 + x[7]*8 + x[8]*9) % 11
    b2 := (x[8]*1 + x[7]*2 + x[6]*3 + x[5]*4 + x[4]*5 + x[3]*6 + x[2]*7 + x[1]*8 + x[0]*9) % 11

    if b1 == 10 {
      b1 = 0
    }
    if b2 == 10 {
      b2 = 0
    }

    if x[9] == b1 && x[10] == b2 {
      resultado[i] = "CPF valido"
    } else {
      resultado[i] = "CPF invalido"
    }
  }

  for _, resultados := range resultado {
    fmt.Println(resultados)
  }
}
