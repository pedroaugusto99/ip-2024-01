package main

import  (
    "fmt"
)
    
func main (){
    var n int 
    
    fmt.Println("Entrada")
    fmt.Scanln(&n)
    
    numeros := make([]int, n,1000)
    for i := 0;i < n;i++{
        fmt.Scan(&numeros[i])
    }
    
    frequencia := make(map[int]int)
    for _, numero := range numeros {
        frequencia[numero]++
    }
  
    fmt.Println("Saída")
    for i, _ := range frequencia {
        fmt.Printf("%d\n",i)
    }
}
