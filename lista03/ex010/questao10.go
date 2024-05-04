package main

import "fmt"

func main() {
    var N int
    fmt.Scan(&N)

    vetor := make([]int, N)
    for i := 0; i < N; i++ {
        fmt.Scan(&vetor[i])
    }

    ui := len(vetor) - 1
    un := vetor[ui]

    frequencia := 0
    for _, nota := range vetor {
        if nota == un {
            frequencia++
        }
    }
    fmt.Printf("Nota %d, %d vezes\n", un, frequencia)

    maiorNota := 0
    indice := 0
    for i := 0; i < N; i++ {
        if vetor[i] > maiorNota {
            maiorNota = vetor[i]
            indice = i
        }
    }
    fmt.Printf("Nota %d, indice %d\n", maiorNota, indice)
}
