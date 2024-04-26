Questão 09

  package main

import "fmt"

func main(){
    var num int
    fmt.Print("Insira um número qualquer:")
    fmt.Scan(&num)
    
    switch {
    case (num < 0):
    fmt.Println("NUMERO INVALIDO!")
    
    case (num == 2):
    fmt.Println("PRIMO")
    case (num == 3):
    fmt.Println("PRIMO")
    case (num % 2 == 0) || (num % 3 == 0):
    fmt.Println("NAO PRIMO")
    case !(num % 2 == 0) || !(num % 3 == 0):
    fmt.Println("PRIMO")
    }
}
