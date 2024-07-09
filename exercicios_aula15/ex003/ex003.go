package main

import "fmt"

func inverterSlice(s []int) {
    for i := 0; i < len(s)/2; i++ {
        s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
    }
}

func main() {
    numeros := []int{1, 2, 3, 4, 5}
    inverterSlice(numeros)
    fmt.Println(numeros)
}
