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




	
Questão 05

package main

import (
	"fmt"
)

// Função para calcular o comprimento do segmento crescente máximo
func maxCrescendoSegment(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	maxLength := 0
	currentLength := 0

	// Iterar pela sequência para encontrar o segmento crescente mais longo
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > numbers[i-1] {
			// Se a sequência continua crescendo, aumente o comprimento atual
			currentLength++
		} else {
			// Se a sequência quebra a ordem crescente, reinicia o comprimento atual
			if currentLength > maxLength {
				maxLength = currentLength
			}
			currentLength = 0 // reiniciar
		}
	}

	// Verifique o segmento atual após o loop
	if currentLength > maxLength {
		maxLength = currentLength
	}

	return maxLength
}

func main() {
	var n int
	fmt.Scan(&n) // Ler o número de elementos na sequência

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i]) // Ler a sequência de números
	}

	maxSegmentLength := maxCrescendoSegment(numbers)

	fmt.Printf("O comprimento do segmento crescente maximo e: %d\n", maxSegmentLength)
}




	
Questão 06

package main

import (
	"fmt"
)

// Função para encontrar a maior subsequência de elementos iguais consecutivos
func maxConsecutiveEqualElements(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	maxLength := 1
	currentLength := 1

	// Percorre a sequência de números
	for i := 1; i < len(numbers); i++ {
		if numbers[i] == numbers[i-1] {
			// Se o número atual é igual ao anterior, aumente a contagem atual
			currentLength++
		} else {
			// Se a sequência é interrompida, compare com a máxima encontrada
			if currentLength > maxLength {
				maxLength = currentLength
			}
			currentLength = 1 // reiniciar a contagem
		}
	}

	// Verifique se a última subsequência é a mais longa
	if currentLength > maxLength {
		maxLength = currentLength
	}

	return maxLength
}

func main() {
	var n int
	fmt.Scan(&n) // Ler o número de elementos na sequência

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i]) // Ler a sequência de números
	}

	maxSubsequence := maxConsecutiveEqualElements(numbers)

	fmt.Printf("A maior subsequência de elementos iguais possui %d elementos.\n", maxSubsequence)
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





	
Questão 11

package main

import (
  "fmt"
  "math"
)

func main() {
  var n, e float64
  fmt.Print("Insira o valor de n: ")
  fmt.Scan(&n)
  fmt.Print("Insira o valor do erro tolerável e: ")
  fmt.Scan(&e)

  
  r := 1.0
  iteracao := 0

  
  for {
    iteracao++
    
    r = (r + n/r) / 2
     
    erro := math.Abs(n - r*r)

    
    fmt.Printf("Iteração %d: Aproximação = %.9f, Erro = %.9f\n", iteracao, r, erro)

  
    if erro <= e {
      break
    }
  }
}




	
Questão 12

package main

import ("fmt"
"math")

func calcularIngressoEsperado(valorIngresso, valorAtual float64) int {
	delta := (valorIngresso - valorAtual) / 0.50
	
	if delta > 0 {
		return 120 + 25*int(delta)
	} else {
		return 120 + 30*int(delta)
	}
}

func calcularLucro(valorIngresso, valorAtual float64, ingressos int) float64 {
	despesasFixas := 200.0
	despesasVariaveis := 0.05 * float64(ingressos)
	receita := valorAtual * float64(ingressos)
	return receita - (despesasFixas + despesasVariaveis)
}

func main() {
	var valorIngresso, valorInicial, valorFinal float64

	fmt.Scan(&valorIngresso, &valorInicial, &valorFinal)

	if valorInicial >= valorFinal {
		fmt.Println("INTERVALO INVALIDO.")
		return
	}

	melhorLucro := -math.MaxFloat64
	var melhorValor float64
	var melhorIngressos int

	for valor := valorInicial; valor <= valorFinal; valor++ {
		ingressos := calcularIngressoEsperado(valorIngresso, valor)
		lucro := calcularLucro(valorIngresso, valor, ingressos)

		fmt.Printf("V: %.2f, N: %d, L: %.2f\n", valor, ingressos, lucro)

		if lucro > melhorLucro {
			melhorLucro = lucro
			melhorValor = valor
			melhorIngressos = ingressos
		}
	}

	if melhorLucro < 0 {
		melhorLucro = 0
		melhorValor = 0
		melhorIngressos = 0
	}

	fmt.Printf("Melhor valor final: %.2f\n", melhorValor)
	fmt.Printf("Lucro: %.2f\n", melhorLucro)
	fmt.Printf("Numero de ingressos: %d\n", melhorIngressos)
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







	
Questão 14

package main

import (
  "fmt"
)

func main() {
  var m, n int

  fmt.Print("Digite o número de linhas (m): ")
  fmt.Scan(&m)
  fmt.Print("Digite o número de colunas (n): ")
  fmt.Scan(&n)

  if m <= 0 || n <= 0 {
    fmt.Println("Dimensões inválidas. Os valores de m e n devem ser maiores que zero.")
    return
  }

  for i := 1; i <= m; i++ {
    
    for j := 1; j <= n; j++ {
       
      if i > j {
        fmt.Printf("(%d-%d)\n—", i, j)
      }
    }
  }
}





	
Questão 15

package main

import (
	"fmt"
	"strconv"
)

// Função para inverter um número inteiro
func reverseInt(num int) int {
	str := strconv.Itoa(num) // Converte o número para string
	revStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		revStr += string(str[i]) // Inverte a string
	}
	revInt, _ := strconv.Atoi(revStr) // Converte a string invertida de volta para inteiro
	return revInt
}

func main() {
	var T int
	fmt.Scan(&T) // Ler o número de casos de teste

	for i := 0; i < T; i++ {
		var A, B int
		fmt.Scan(&A, &B) // Ler os dois números de três dígitos

		revA := reverseInt(A) // Inverter A
		revB := reverseInt(B) // Inverter B

		if revA > revB {
			fmt.Println(revA) // Se revA for maior, imprime revA
		} else {
			fmt.Println(revB) // Se revB for maior, imprime revB
		}
	}
}




	
Questão 16

package main

import (
	"fmt"
	"sort"
)

type Cateto struct {
	a, b int
}

// Encontrar pares de catetos para uma dada hipotenusa
func encontrarCatetos(hip int) []Cateto {
	hip2 := hip * hip
	catetos := []Cateto{}
	for a := 1; a < hip; a++ { // a deve ser menor que a hipotenusa
		for b := a; b < hip; b++ { // b deve ser maior ou igual a a, mas menor que a hipotenusa
			if a*a+b*b == hip2 { // Teste do Teorema de Pitágoras
				catetos = append(catetos, Cateto{a, b})
			}
		}
	}
	return catetos
}

func main() {
	var n int
	fmt.Scan(&n) // Ler o valor de n

	for h := 1; h <= n; h++ { // Iterar sobre as hipotenusas de 1 a n
		catetos := encontrarCatetos(h) // Encontrar catetos para uma hipotenusa
		if len(catetos) > 0 { // Se existirem catetos válidos para esta hipotenusa
			sort.Slice(catetos, func(i, j int) bool { // Ordenar por a, depois por b
				return catetos[i].a < catetos[j].a
			})
			for _, c := range catetos { // Imprimir todos os pares de catetos
				fmt.Printf("hipotenusa = %d, catetos %d e %d\n", h, c.a, c.b)
			}
		}
	}
}





	
Questão 17

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		lucroMenorQue10, lucroEntre10E20, lucroMaiorQue20 int
		maiorLucro, totalCompras, totalVendas             float64
		codigoMaiorLucro, codigoMaiorVendas               int
	)

	for {
		var codigo, numVendas int
		var precoCompra, precoVenda float64

		_, err := fmt.Scan(&codigo, &precoCompra, &precoVenda, &numVendas)
		if err != nil {
			break
		}

		lucro := ((precoVenda - precoCompra) / precoCompra) * 100

		switch {
		case lucro < 10:
			lucroMenorQue10++
		case lucro <= 20:
			lucroEntre10E20++
		default:
			lucroMaiorQue20++
		}

		if lucro > maiorLucro {
			maiorLucro = lucro
			codigoMaiorLucro = codigo
		}

		if numVendas > codigoMaiorVendas {
			codigoMaiorVendas = numVendas
		}

		totalCompras += precoCompra * float64(numVendas)
		totalVendas += precoVenda * float64(numVendas)
	}

	lucroTotal := totalVendas - totalCompras
	lucroPercentualTotal := (lucroTotal / totalCompras) * 100

	fmt.Printf("Quantidade de mercadorias que geraram lucro menor que 10%%: %d\n", lucroMenorQue10)
	fmt.Printf("Quantidade de mercadorias que geraram lucro entre 10%% e 20%%: %d\n", lucroEntre10E20)
	fmt.Printf("Quantidade de mercadorias que geraram lucro maior que 20%%: %d\n", lucroMaiorQue20)
	fmt.Printf("Codigo da mercadoria que gerou maior lucro: %d\n", codigoMaiorLucro)
	fmt.Printf("Codigo da mercadoria mais vendida: %d\n", codigoMaiorVendas)
	fmt.Printf("Valor total de compras: %.2f, valor total de vendas: %.2f e percentual de lucro total: %.2f%%\n", totalCompras, totalVendas, lucroPercentualTotal)
}




	

Questão 18

package main

import (
  "fmt"
  "math"
)

// Função para verificar se um número é primo
func isPrimo(n int) bool {
  if n <= 1 {
    return false
  }
  for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
    if n%i == 0 {
      return false
    }
  }
  return true
}

// Função para obter o próximo número primo após um dado número
func proximoPrimo(n int) int {
  for i := n + 1; ; i++ {
    if isPrimo(i) {
      return i
    }
  }
}

// Função para calcular o MMC usando fatoração
func calcularMMC(a, b, c int) int {
  n1, n2, n3 := a, b, c
  mmc := 1
  primo := 2

  fmt.Printf("%2d %2d %2d :", n1, n2, n3)

  // Enquanto pelo menos um dos números não for 1
  for n1 != 1 || n2 != 1 || n3 != 1 {
    dividiu := false

    if n1%primo == 0 {
      n1 /= primo
      dividiu = true
    }

    if n2%primo == 0 {
      n2 /= primo
      dividiu = true
    }

    if n3%primo == 0 {
      n3 /= primo
      dividiu = true
    }

    if dividiu {
      fmt.Printf(" %2d\n", primo) // Exibe o primo usado para fatorar
      fmt.Printf("%2d %2d %2d :", n1, n2, n3)
      mmc *= primo
    } else {
      primo = proximoPrimo(primo) // Obter o próximo primo
    }
  }

  return mmc
}

func main() {
  var a, b, c int

  // Leitura dos três números
  fmt.Println("Digite três números inteiros:")
  fmt.Scan(&a, &b, &c)

  // Calcular o MMC usando a fatoração
  fmt.Println("Processo de fatoração para calcular o MMC:")
  mmc := calcularMMC(a, b, c)

  fmt.Printf("\nMMC de %d, %d e %d é: %d\n", a, b, c, mmc)
}




	
Questão 19

package main

import "fmt"

func main() {
  var m int
  
  fmt.Print("Digite um número inteiro maior que zero: ")
  fmt.Scanln(&m)
  
  if m <= 0 {
    fmt.Println("Dígito Inválido")
  }
  
  for k := 1; k <= m; k++ {
    soma := 0
    inicio := k*k - k + 1 // O primeiro ímpar da sequência
    sequencia := make([]int, k) // Para armazenar a sequência de ímpares

    for i := 0; i < k; i++ {
      sequencia[i] = inicio + 2*i // Gera os números ímpares consecutivos
      soma += sequencia[i] // Calcula a soma
    }

    // Formatar a saída
    fmt.Printf("%d×%d×%d = ", k, k, k)
    for i, v := range sequencia {
      if i == len(sequencia)-1 {
        fmt.Printf("%d\n", v) // Sem + no final
      } else {
        fmt.Printf("%d+", v) // Com + para intermediários
      }
    }
}
}






	
Questão 20

package main

import (
	"fmt"
	"strings"
)

func encontrarDivisores(n int) []int {
	var divisores []int
	for i := 1; i <= n/2; i++ { // Divisores vão até n/2
		if n%i == 0 {
			divisores = append(divisores, i)
		}
	}
	return divisores
}

func somaDivisores(divisores []int) int {
	soma := 0
	for _, divisor := range divisores {
		soma += divisor
	}
	return soma
}

func main() {
	
	var n int
	fmt.Scan(&n)

	
	divisores := encontrarDivisores(n)

	
	somaDiv := somaDivisores(divisores)

	
	var mensagem string
	if somaDiv == n {
		mensagem = "NUMERO PERFEITO"
	} else {
		mensagem = "NUMERO NAO E PERFEITO"
	}

	
	divisoresStr := make([]string, len(divisores))
	for i, divisor := range divisores {
		divisoresStr[i] = fmt.Sprintf("%d", divisor)
	}

	divisoresFormatados := strings.Join(divisoresStr, " + ")

	fmt.Printf("%d = %s = %d (%s)\n", n, divisoresFormatados, somaDiv, mensagem)
}





	
Questão 21
	
package main

import (
  "fmt"
)

func main() {
  for {
    // Ler o tamanho da próxima sequência
    var n int
    if _, err := fmt.Scan(&n); err != nil {
      // Se erro ou EOF, terminar o loop
      break
    }

    // Se o tamanho for zero, termina
    if n == 0 {
      break
    }

    // Ler a sequência de números
    numbers := make([]float64, n)
    for i := 0; i < n; i++ {
      _, err := fmt.Scan(&numbers[i])
      if err != nil {
        break
      }
    }

    // Verificar se está ordenada
    isOrdered := true
    for i := 0; i < len(numbers); i++ {
      if numbers[i] >= numbers[i+1] {
        isOrdered = false
        break
      }
    }

    // Imprimir a resposta conforme instruído
    if isOrdered {
      fmt.Println("ORDENADA")
    } else  {
      fmt.Println("DESORDENADA")
    } 
  }
}






	
Questão 22

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Função para calcular o Máximo Divisor Comum (MDC) usando o algoritmo de Euclides
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// Função para converter um número decimal para uma fração simplificada
func decimalParaFracao(n float64) (int, int) {
	// Convertendo o número para string para analisar a parte decimal
	nStr := fmt.Sprintf("%.10f", n) // Limitar a 10 casas decimais
	nStr = strings.TrimRight(nStr, "0") // Remover zeros à direita
	parts := strings.Split(nStr, ".")

	// Parte inteira do número
	integral, _ := strconv.Atoi(parts[0])

	// Parte decimal
	if len(parts) < 2 {
		// Não há parte decimal, então é uma fração com denominador 1
		return integral, 1
	}

	decimalPart := parts[1]
	decimalDigits := len(decimalPart)

	// Denominador é 10^decimalDigits
	denominator := int(math.Pow(10, float64(decimalDigits)))
	// Numerador é a parte inteira vezes o denominador mais a parte decimal
	decimalInt, err := strconv.Atoi(decimalPart) // Trata o erro ao converter a parte decimal para número
	if err != nil {
		fmt.Println("Erro ao converter parte decimal para número:", err)
		return 0, 0
	}
	numerator := integral*denominator + decimalInt

	// Obter o MDC para simplificar a fração
	mdc := gcd(numerator, denominator)

	// Retornar a fração simplificada
	return numerator / mdc, denominator / mdc
}

func main() {
	var n float64
	fmt.Scan(&n)

	num, den := decimalParaFracao(n)

	fmt.Printf("%d/%d\n", num, den)
}







	
Questão 23

package main

import (
	"fmt"
)

// Função para encontrar todos os divisores de um número
func encontrarDivisores(n int) []int {
	var divisores []int
	for i := 1; i <= n/2; i++ { // Divisores são todos os números até n/2
		if n%i == 0 {
			divisores = append(divisores, i)
		}
	}
	return divisores
}

// Função para calcular a soma dos divisores
func somaDivisores(divisores []int) int {
	soma := 0
	for _, divisor := range divisores {
		soma += divisor
	}
	return soma
}

func main() {
	var n int
	fmt.Scan(&n)

	encontrados := 0
	paresAmigos := make(map[int]bool) // Mapeamento para evitar duplicidades
	num := 1

	for encontrados < n {
		divisores := encontrarDivisores(num)
		soma := somaDivisores(divisores)

		// Verificar se são amigos e se não são repetidos
		if soma > num { // Garante que o menor número vem primeiro
			divisoresSoma := encontrarDivisores(soma)
			somaDoOutro := somaDivisores(divisoresSoma)
			if somaDoOutro == num && !paresAmigos[num] && !paresAmigos[soma] {
				fmt.Printf("(%d,%d)\n", num, soma)
				paresAmigos[num] = true
				paresAmigos[soma] = true
				encontrados++
			}
		}

		num++ // Incrementa o número para buscar o próximo par de números amigos
	}
}



	

	

Questão 24

package main

import (
	"fmt"
	"math"
)

// Função para calcular o fatorial de um número
func fatorial(n int) int {
	resultado := 1
	for i := 2; i <= n; i++ {
		resultado *= i
	}
	return resultado
}

// Função para calcular o cosseno usando a série de Taylor
func cosTaylor(x float64, n int) float64 {
	cos := 0.0
	for i := 0; i <= n; i++ { // N é o número de termos
		termo := math.Pow(-1, float64(i)) * math.Pow(x, float64(2*i)) / float64(fatorial(2*i))
		cos += termo
	}
	return cos
}

func main() {
	var x float64
	var n int
	fmt.Scan(&x, &n)

	// Calcular o cosseno de x usando a série de Taylor com N termos
	cos := cosTaylor(x, n)

	// Imprimir o resultado no formato especificado
	fmt.Printf("cos(%.2f) = %.6f\n", x, cos)
}






	
Questão 25

package main

import (
	"fmt"
	"math"
)

// Função para calcular o fatorial de um número
func fatorial(n int) int {
	resultado := 1
	for i := 2; i <= n; i++ {
		resultado *= i
	}
	return resultado
}

// Função para calcular o valor de e^x usando a série de Taylor até N termos
func euler(x float64, n int) float64 {
	euler := 0.0
	for i := 0; i <= n; i++ {
		termo := math.Pow(x, float64(i)) / float64(fatorial(i))
		euler += termo
	}
	return euler
}

func main() {
	var x float64
	var n int
	fmt.Scan(&x, &n)

	// Calcular e^x usando a série de Taylor até N termos
	eulerValue := euler(x, n)

	// Imprimir o resultado com formato especificado
	fmt.Printf("e^%.2f = %.6f\n", x, eulerValue)
}






	
Questão 26

package main

import (
	"fmt"
	"math"
)

// Função para calcular o fatorial de um número
func fatorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * fatorial(n-1)
}

// Função para calcular o seno usando a série de Taylor
func senoTaylor(x float64, n int) float64 {
	seno := 0.0
	for i := 0; i <= n; i++ {
		termo := math.Pow(-1, float64(i)) * math.Pow(x, float64(2*i+1)) / float64(fatorial(2*i+1))
		seno += termo
	}
	return seno
}

func main() {
	var x float64
	var n int
	fmt.Scan(&x, &n)

	// Calcular o seno de x usando a série de Taylor com N termos
	sin := senoTaylor(x, n)

	// Imprimir o resultado no formato especificado
	fmt.Printf("seno(%.2f) = %.6f\n", x, sin)
}






	
Questão 27

package main

import (
	"fmt"
)

// Função para verificar se um número é primo
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Função para fatorar um número em fatores primos
func primeFactorization(n int) []int {
	factors := []int{}
	if n <= 1 {
		return factors
	}

	// Divisão por 2 até que não seja mais divisível
	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}

	// Testar todos os números ímpares de 3 até √n
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}

	// Se sobrou algum número maior que 1, é um fator primo
	if n > 1 {
		factors = append(factors, n)
	}

	return factors
}

func main() {
	var n int
	fmt.Scan(&n)

	if n <= 1 {
		fmt.Printf("Fatoracao nao e possivel para o numero %d!\n", n)
		return
	}

	factors := primeFactorization(n)

	if len(factors) == 0 {
		fmt.Printf("Fatoracao nao e possivel para o numero %d!\n", n)
	} else {
		fmt.Printf("%d = ", n)
		for i, factor := range factors {
			if i > 0 {
				fmt.Print(" x ")
			}
			fmt.Print(factor)
		}
		fmt.Println() // para a quebra de linha ao final
	}
}

