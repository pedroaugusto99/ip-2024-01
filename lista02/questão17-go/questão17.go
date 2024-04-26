Questão 17

package main

import "fmt"

func main() {
  var (
    lucroMenorQue10, lucroEntre10E20, lucroMaiorQue20 int
    maiorLucro, totalCompras, totalVendas             float64
    codigoMaiorLucro, codigoMercadoriaMaisVendida     int
    maiorVendas                                       int
  )

  for {
    var codigo, numVendas int
    var precoCompra, precoVenda float64

    _, err := fmt.Scan(&codigo, &precoCompra, &precoVenda, &numVendas)
    if err != nil {
      break
    }

    lucro := ((precoVenda - precoCompra) / precoCompra) * 100

    switch {
    case lucro < 10:
      lucroMenorQue10++
    case lucro <= 20:
      lucroEntre10E20++
    default:
      lucroMaiorQue20++
    }

    if lucro > maiorLucro {
      maiorLucro = lucro
      codigoMaiorLucro = codigo
    }

    if numVendas > maiorVendas {
      maiorVendas = numVendas
      codigoMercadoriaMaisVendida = codigo
    }

    totalCompras += precoCompra * float64(numVendas)
    totalVendas += precoVenda * float64(numVendas)
  }

  lucroTotal := totalVendas - totalCompras
  lucroPercentualTotal := (lucroTotal / totalCompras) * 100

  fmt.Printf("Quantidade de mercadorias que geraram lucro menor que 10%%: %d\n", lucroMenorQue10)
  fmt.Printf("Quantidade de mercadorias que geraram lucro entre 10%% e 20%%: %d\n", lucroEntre10E20)
  fmt.Printf("Quantidade de mercadorias que geraram lucro maior que 20%%: %d\n", lucroMaiorQue20)
  fmt.Printf("Código da mercadoria que gerou maior lucro: %d\n", codigoMaiorLucro)
  fmt.Printf("Código da mercadoria mais vendida: %d\n", codigoMercadoriaMaisVendida)
  fmt.Printf("Valor total de compras: %.2f, valor total de vendas: %.2f e percentual de lucro total: %.2f%%\n", totalCompras, totalVendas, lucroPercentualTotal)
}
