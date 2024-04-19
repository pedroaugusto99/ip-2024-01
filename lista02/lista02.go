Questão 01 ||

package main

import "fmt"

func main() {
         var p1,p2,p3,p4,p5,p6,p7,p8 float64
         var l1,l2,l3,l4,l5 float64
         var m int
         var presenca, mp, ml, p, nt, nf float64
         
    fmt.Print("Digite a matrícula do aluno: ")
    fmt.Scanln(&m)
    fmt.Print("Digite a nota das 8 provas: ")
    fmt.Scanln(&p1,&p2,&p3,&p4,&p5,&p6,&p7,&p8)
    fmt.Print("Digitea nota dos 5 exercícios: ")
    fmt.Scanln(&l1,&l2,&l3,&l4,&l5)
    fmt.Print("Digite a nota do trabalho final: ")
    fmt.Scanln(&nt)
    fmt.Print("Digite a presença: ")
    fmt.Scanln(&presenca)
    
    mp = (p1+p2+p3+p4+p5+p6+p7+p8)/8
    ml = (l1+l2+l3+l4+l5)/5
    p = presenca/128
    
    nf = 0.7*mp + 0.15*ml + 0.15*nt
    
    switch {
        case (nf >= 6) && (p >= 0.75):
        fmt.Printf("Matricula: %v, Nota Final: %.2f, Situação Fonal: APROVADO",m,nf)
        
        case (nf < 6) && (p >= 0.75):
        fmt.Printf("Matricula: %v, Nota Final: %.2f, Situação Final: REPROVADO POR NOTA",m,nf)
        
        case (nf >= 6) && (p < 0.75):
        fmt.Printf("Matricula: %v, Nota Final: %.2f, Situação Final: REPROVADO POR FREQUENCIA",m,nf)
        
        case (nf < 6) && (p < 0.75):
       fmt.Printf("Matricula: %v, Nota Final: %.2f, Situação Final: REPROVADO POR NOTA E POR FREEQUENCIA",m,nf)
    }
         
  Questão 02

package main

import (
  "fmt"
)

func main() {
  var popuA, popuB float64
  
  fmt.Print("População País A: ")
  fmt.Scanln(&popuA)
  fmt.Print("População País B: ")
  fmt.Scanln(&popuB)

  anos := 0
  
    for popuA <= popuB {
    popuA = popuA + popuA * 0.03
    popuB = popuB + popuB * 0.015
    anos++
  }

  fmt.Printf("ANOS = %d\n", anos)
}
         
Questão 03

package main

import "fmt"

func main() {
  var n int
  fmt.Print("Digite um número qualquer:")
  fmt.Scanln(&n)

  fatorial := 1
  for i := 1; i <= n; i++ {
    fatorial = fatorial * i
  }

  fmt.Printf("%d! = %d\n", n, fatorial)
}

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

Questão 07

package main

import "fmt"

func main() {
    var num int
    var par, somapar, impar, somaimpar float64

    fmt.Println("Insira uma sequência de    números inteiros diferentes de zero (Use o 0 para encerrar):")
 
    for {
        fmt.Scan(&num)
        
        if num == 0 {
            break
        }  
        
        if num%2 == 0{
            
            somapar = somapar + float64(num)
            par++
            
        } else {
            
            somaimpar = somaimpar + float64(num)
            impar++
            }
             
    }
    
    var mediapar, mediaimpar float64

    if par > 0 {
    mediapar = somapar/ par
    }
    if impar > 0 {
    mediaimpar = somaimpar / impar
    }

    fmt.Printf("MEDIA PAR = %.2f\n", mediapar)
    fmt.Printf("MEDIA IMPAR = %.2f\n", mediaimpar)
}

Questão 08

package main

import (
  "fmt"
)

func main() {
  var x, final int
  fmt.Print("Digite a quantidade de times: ")
  fmt.Scan(&x)

  if x < 2 {
    fmt.Println("Campeonato inválido!")
    return
  }
  
  final = 1
  
  for i := 1; i <= x; i++ {
    for j := i + 1; j <= x; j++ {
      fmt.Printf("Final %d: Time%d X Time%d\n", final, i, j)
      final++
    }
  }
}

Questão 10

package main

import "fmt"

func main(){
  
  var (
    matricula, horas int
    salarioporhora, salario float64
  )
  
 
  
  for {
    
  fmt.Println("Digite a matrícula, horas trabalhadas e salário por hora:")
  
  fmt.Scan(&matricula, &horas, &salarioporhora)
  
  
  
  if matricula == 0 {
    break
  }
  
  salario = float64(horas) * salarioporhora
  fmt.Printf("%d %.2f\n", matricula, salario)
  } 
}


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
