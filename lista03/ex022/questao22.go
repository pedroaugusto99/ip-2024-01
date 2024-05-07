package main

import "fmt"

func encontrarMaiorNumero(numero string, d int) string {
  n := len(numero)
  if d >= n {
    return numero
  }

  digitos := make([]byte, 0, n)
  aRemover := n - d

  for i := 0; i < n; i++ {
    digitoAtual := numero[i]

    for len(digitos) > 0 && digitos[len(digitos)-1] < digitoAtual && aRemover > 0 {
      digitos =digitos[:len(digitos)-1]
      aRemover--
    }

    digitos = append(digitos, digitoAtual)
  }

  return string(digitos[:d])
}

func main() {
    var resultados []string
  var n, d int
  var numero string

  for {
    fmt.Scanf("%d %d", &n, &d)
    if n == 0 && d == 0 {
      break
    }
    fmt.Scan(&numero)

    maiorNumero := encontrarMaiorNumero(numero, d)
    resultados = append(resultados,maiorNumero)
 
  }
  for _, value := range resultados {
    fmt.Println(value)
  }
}
