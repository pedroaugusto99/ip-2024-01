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
