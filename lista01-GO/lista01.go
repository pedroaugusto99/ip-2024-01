Questão 01

package main

import "fmt"

func main() {
	var nota1, nota2, nota3 float64

  fmt.Println("Digite as 3 notas do aluno separadas entre si por um caractere de espaço:") 
	fmt.Scan(&nota1, &nota2, &nota3)

	media := (nota1 + nota2 + nota3) / 3

	var mensagem string
	if media >= 6 {
		mensagem = "APROVADO"
	} else {
		mensagem = "REPROVADO"
	}

	fmt.Printf("MEDIA = %.2f\n%s\n", media, mensagem)
}


Questão 02

package main

import "fmt"

func main() {
    var njogos int
    fmt.Println("Número de jogos:")
    fmt.Scan(&njogos)

    for i := 1; i <= njogos; i++ {
        var totalPessoas, popular, geral, arquibancada, cadeiras float64
        fmt.Scan(&totalPessoas, &popular, &geral, &arquibancada, &cadeiras)

        renda := (popular/100)*totalPessoas*1 + (geral/100)*totalPessoas*5 + (arquibancada/100)*totalPessoas*10 + (cadeiras/100)*totalPessoas*20

        fmt.Printf("A RENDA DO JOGO N. %d E = %.2f\n", i, renda)
    }
}

Questão 03

package main

import "fmt"

func main() {
    var n1, n2, n3 int
    fmt.Println("Digite três números inteiros separados por espaço:")
    fmt.Scan(&n1, &n2, &n3)
 
    if n1 < 0 || n1 > 9 || n2 < 0 || n2 > 9 || n3 < 0 || n3 > 9 {
        fmt.Println("DIGITO INVALIDO")
        return     
   }
   
   var conc, concquadrado int
   
   conc = (n1*100) + (n2*10) + n3
   
   concquadrado = conc * conc
   
   fmt.Println(conc, concquadrado)
}

Questão 04

package main

import "fmt"

func main() {
    
    var sm, kW float64
    fmt.Println("Digite o valor do salário mínimo: R$ ")
    fmt.Scan(&sm)
    fmt.Println("Digite a quantidade de kW gasta pela residência: ")
    fmt.Scan(&kW)
    
    valorporKW := sm * 0.7 / 100
    
    cc := valorporKW * kW
    
    cd := cc * 0.9
    
    fmt.Printf("CUSTO POR KW: R$ %.2f\n", valorporKW)
    fmt.Printf("CUSTO DO CONSUMO: R$ %.2f\n", cc)
    fmt.Printf("CUSTO COM DESCONTO: R$ %.2f\n", cd)
}

Questão 05

package main

import "fmt"

func main() {
  var contac int
  var consumo, valor float64
  var tipoc string

  fmt.Print("Digite o número da conta do cliente:")
  fmt.Scanln(&contac)

  fmt.Print("Digite o consumo de água em metros cúbicos: ")
  fmt.Scanln(&consumo)

  fmt.Print("Digite o tipo do consumidor (R para residencial, C para comercial ou I para industrial): ")
  fmt.Scanln(&tipoc)

  if tipoc == "R" {
    valor = 5 + 0.05*consumo
  } else if tipoc == "C" {
    if consumo <= 80 {
      valor = 500
    } else {
      valor = 500 + 0.25*(consumo-80)
    }
  } else if tipoc == "I" {
    if consumo <= 100 {
      valor = 800
    } else {
      valor = 800 + 0.04*(consumo-100)
    }
  }

  fmt.Println("CONTA =", contac)
  fmt.Printf("VALOR DA CONTA = %.2f\n", valor)
}



Questão 06

package main

import "fmt"

func main() {
    var ntemp int
    fmt.Println("Número de temperaturas a serem convertidas:")
    fmt.Scan(&ntemp)

    for i := 1; i <= ntemp; i++ {
        var c,f float64
        fmt.Println("Temperatura em Fahrenheit(F°):")
        fmt.Scan(&f)
        c = (f - 32)/1.8
        x := "CELSIUS"
        fmt.Printf("%v FAHRENHEIT EQUIVALE A %.2f %v\n",f,c,x)
    }
}

Questão 07

package main

import "fmt"

func main() {
    var c, f, polg, mm float64
    fmt.Println("Digite a temperatura em Fahrenheit(F°):")
    fmt.Scan(&f)
    fmt.Println("Digite a quantidade de chuva em polegadas:")
    fmt.Scan(&polg)
    
    c = (f - 32) / 1.8
    mm = polg * 25.4
    
    fmt.Printf("O VALOR EM CELSIUS = %.2f \n",c)
    fmt.Printf("A QUANTIDADE DE CHUVA E = %.2f",mm)
    
    
}

Questão 08

package main

import "fmt"

func main() {
    var custo, raio, altura, arealateral, areadabase float64
    fmt.Println("Digite a altura em metros:")
    fmt.Scan(&altura)
    fmt.Println("Digite o raio em metros:")
    fmt.Scan(&raio)
    
    areadabase = (3.14 * raio * raio) * 2
    arealateral = 2*3.14*raio*altura
    
    custo = (arealateral + areadabase) * 100
    
    fmt.Printf("O VALOR DO CUSTO E R$ %.2f \n",custo)
      
}

Questão 09

package main

import "fmt"

func main() {
    var a, b, c, delta float64
    fmt.Println("A:")
    fmt.Scan(&a)
    fmt.Println("B:")
    fmt.Scan(&b)
    fmt.Println("C:")
    fmt.Scan(&c)
    
    delta = b*b - 4 * a * c
    
    
    fmt.Printf("O VALOR DE DELTA E %.2f \n",delta)
      
}

Questão 10

package main

import "fmt"

func main() {
    var a, b, c, d float64

    fmt.Print("Digite A: ")
    fmt.Scanln(&a)
    fmt.Print("Digite B: ")
    fmt.Scanln(&b)
    fmt.Print("Digite C: ")
    fmt.Scanln(&c)
    fmt.Print("Digite D: ")
    fmt.Scanln(&d)

    determinante := (a*d - b*c)

    fmt.Printf("O VALOR DO DETERMINANTE E = %.2f\n", determinante)
}


Questão 11

package main

import "fmt"

func main() {
    var x int
    fmt.Println("Insira um número:")
    fmt.Scan(&x)
    
    if (x % 3 == 0) && (x % 5 == 0) {
    fmt.Println("O NÚMERO É DIVISÍVEL")
    } else {
    fmt.Println("O NÚMERO NÃO É DIVISÍVEL")
    }
      
}

Questão 12

package main

import "fmt"

func main() {
  var horas float64
  var valor float64

  fmt.Print("Digite o número de horas: ")
  fmt.Scan(&horas)

  switch {
  case int(horas)%3 == 0:
    valor = (horas / 3) * 10
    fmt.Println("Valor:", valor)
  case int(horas)%3 !=0:
    x := int(horas)% 3
    y := int(horas) - int(x)
    z := y/3
    a := x*5
    b := z*10
    valor = float64(a) + float64(b)
    fmt.Printf("O VALOR A PAGAR E = %.2f \n", valor)
  }
}



Questão 13

package main

import "fmt"

func main() {
    var x float64
    fmt.Println("Insira uma nota de 0 a 10:")
    fmt.Scan(&x)
    
    switch {
      case x >= 9:
        fmt.Printf("Nota = %v CONCEITO = A\n", x)
      case (x < 9) && (x >= 7.5):
        fmt.Printf("Nota = %v CONCEITO = B\n", x)
      case (x < 7.5) && (x >= 6):
        fmt.Printf("Nota = %v CONCEITO = C\n", x)
      case (x < 6):
        fmt.Printf("Nota = %v CONCEITO = D\n", x)
    }
      
}

Questão 14

package main

import "fmt"

func main() {
     
     var h, a float64
     
     fmt.Print("Digite a altura da pirâmide: ")
     fmt.Scanln(&h)
     fmt.Print("Digite a aresta da base: ")
     fmt.Scanln(&a)
    
      var ab, volume float64
       ab = 3 * (a*a) * 0.866025404 
       
       volume = (ab * h)/3
       
       fmt.Printf("O VOLUME DA PIR MIDE E = %.2f METROS CÚBICOS\n", volume)       
}

Questão 15

package main

import "fmt"

func main() {
  var N int
  fmt.Println("Digite um número aleatório entre 5 e 2000: ")
  fmt.Scan(&N)

  for x := 2; x <= N; x = x+2 {
    quadrado := x * x
    fmt.Printf("%d^2 = %d\n", x, quadrado)
  }
}



Questão 16

package main

import "fmt"

func main() {
    var x float64
    fmt.Println("Insira o salário(R$):")
    fmt.Scan(&x)
    
    switch {
      case x <= 300 :
        x = x*1.5
        fmt.Printf("SALÁRIO COM REAJUSTE = %.2f ", x)
      case x > 300:
        x = x*1.3
        fmt.Printf("SALÁRIO COM REAJUSTE = %.2f ", x)
    }

Questão 17

package main

import "fmt"

func main() {
  var y, x int
  fmt.Print("Digite o valor de x:")
  fmt.Scanln(&x)
  fmt.Print("Digite o valor de y: ")
  fmt.Scanln(&y)
  
  switch {
    
    case x % 2 == 0:

    for i := 0; i < y;i++ {
    fmt.Print(x," ")
    x += 2}
    
    case x % 2 != 0:
      fmt.Println("O NÚMERO INSERIDO(X) NÃO É PAR")

  
  }
}
      
}

Questão 18

package main

import "fmt"

func main() {
        var a1, razao, num, soma int
        fmt.Print("Digite o valor de a1: ")
        fmt.Scanln(&a1)
        fmt.Print("Digite o valor da razão: ")
        fmt.Scanln(&razao)
        fmt.Print("Digite o número de elementos: ")
        fmt.Scanln(&num)
        
        soma = 0
    for x := 0; x < num; x++ {
        soma = soma + a1 + x*razao
        }
        
        fmt.Println(soma)
  
}

Questão 19

package main

import "fmt"

func main() {
  var x int
  var soma float64
  fmt.Print("Digite um número inteiro e positivo:")
  fmt.Scanln(&x)
  
  switch {
    
    case (x > 0):
      
     
   for i := 1; i <= x; i++ {
     soma += 1.0 / float64(i)
}
      fmt.Printf("%.6f\n", soma)
      case (x < 0):
      fmt.Println("NÚMERO INVÁLIDO!")
}
}

Questão 20

package main

import "fmt"

func main() {
    var h, m, s, tempo int
    fmt.Println("Insira o tempo em horas:")
    fmt.Scan(&h)
    fmt.Println("Insira o tempo em minutos:")
    fmt.Scan(&m)
    fmt.Println("Insira o tempo em segundos:")
    fmt.Scan(&s)
    
    tempo = (h*3600) + (m*60) + s
    
    fmt.Println("O TEMPO EM SEGUNDOS E =",tempo)
      
}
