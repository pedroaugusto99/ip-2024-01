package main

import "fmt"

func Soma(x []float64) float64 {
    if len(x) == 0 {
        return 0
    }
    
    return x[0] + Soma(x[1:])
}

func main() {
   x := []float64{2,3,4,5}
   
   fmt.Println(Soma(x))
}
