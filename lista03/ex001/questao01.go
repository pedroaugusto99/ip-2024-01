package main

import "fmt"

func main() {
   
    var n int
    
    fmt.Println("Entrada")
    fmt.Scanln(&n)

    numeros := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&numeros[i])
    }

    mapaNumeros := make(map[int]bool)
    for _, v := range numeros {
        mapaNumeros[v] = true
    }

    var m int
    fmt.Scanln(&m)

    resultados := make([]string, m)

    for i := 0; i < m; i++ {
        var busca int
        fmt.Scanln(&busca)

        if mapaNumeros[busca] {
            resultados[i] = "ACHEI"
        } else {
            resultados[i] = "NAO ACHEI"
        }
    }
    fmt.Println("Saída\n")
    
    for _, resultado := range resultados {
        fmt.Println(resultado)
    }
}
