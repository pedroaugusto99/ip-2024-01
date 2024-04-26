Questão 04

package main

import "fmt"

func main() {
  
  var n, i, K, s float64
  
  fmt.Print("Digite um número inteiro de 0 a 9:  ")
  fmt.Scanln(&n)
  fmt.Print("Digite o número inicial(i):  ")
  fmt.Scanln(&i)
  fmt.Print("Insira o número de valores (K):  ")
  fmt.Scanln(&K)
  fmt.Print("Insira o valor do incremento (s):  ")
  fmt.Scanln(&s)
  
  fmt.Print("\n")
  
  var  soma1, soma2, sub1, sub2, multi1, multi2, div1, div2 float64
  var increm float64
  
  increm = i + (i * s) 
    
  // Soma
  
  soma1 = n + i
  soma2 = n + increm
  
  fmt.Println("Tabuada de soma:")
  fmt.Printf("%.2f + %.2f = %.2f\n", n, i, soma1)
  fmt.Printf("%.2f + %.2f = %.2f\n", n, i, soma2)
  
  fmt.Println("\n")
  
  // Subtração
  
  sub1 = n - i
  sub2 = n - increm
  
  fmt.Println("Tabuada de subtração:")
  fmt.Printf("%.2f - %.2f = %.2f\n", n, i, sub1)
  fmt.Printf("%.2f - %.2f = %.2f\n", n, i, sub2)
  
  fmt.Println("\n")
  
  //Multiplicação
  
  multi1 = n * i
  multi2 = n * increm
  
  fmt.Println("Tabuada de multiplicação:")
  fmt.Printf("%.2f × %.2f = %.2f\n", n, i, multi1)
  fmt.Printf("%.2f × %.2f = %.2f\n", n, i, multi2)
  
  fmt.Println("\n")
   
  //Divisão
  
  div1 = n / i
  div2 = n / increm
  
  fmt.Println("Tabuada de divisão:")
  fmt.Printf("%.2f / %.2f = %.2f\n", n, i, div1)
  fmt.Printf("%.2f / %.2f = %.2f\n", n, i, div2)
}
