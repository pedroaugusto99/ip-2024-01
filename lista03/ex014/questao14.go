package main

import (
    "fmt"
    "sort"
    )

func main() {
 var n int
 
 fmt.Scan(&n)
 
 vetor := make([]int,n)
 for i := 0; i < n;i++{
     fmt.Scan(&vetor[i])
 }
 
 sort.Ints(vetor)
 
 var mediana float64
 
 if n%2 == 0{
  mediana = float64(vetor[n/2-1] + vetor[n/2])/2
 } else {
     mediana = float64(vetor[n/2])
 }
 
 fmt.Printf("%.2f\n", mediana)
}
