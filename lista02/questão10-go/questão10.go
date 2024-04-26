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
