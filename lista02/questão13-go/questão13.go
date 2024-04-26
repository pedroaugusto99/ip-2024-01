Questão 13

  package main

  import "fmt"

  func main(){
    var num, x int
    fmt.Print("Insira um número qualquer:")
    fmt.Scanln(&num)
    
    x = ((2*num)*32 + num*32) - num
    fmt.Println(x)
}
