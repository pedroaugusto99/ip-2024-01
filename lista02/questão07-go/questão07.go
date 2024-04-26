Questão 07
		
package main

import "fmt"

func main() {
    var num int
    var Qtdpar, somapar, Qtdimpar, somaimpar float64

    fmt.Println("Insira uma sequência de números inteiros diferentes de zero (Use o 0 para encerrar):")
 
    Qtdpar = 0
    Qtdimpar = 0
 
 
    for {
        fmt.Scan(&num)
        
        if num == 0 {
            break
        }  
        
        if num%2 == 0{
            
            somapar = somapar + float64(num)
            Qtdpar++
            
        } else {
            
            somaimpar = somaimpar + float64(num)
            Qtdimpar++
            }
             
    }
    
    var mediapar, mediaimpar float64

    if Qtdpar > 0 {
    mediapar = somapar/ Qtdpar
    }
    if Qtdimpar > 0 {
    mediaimpar = somaimpar / Qtdimpar
    }

    fmt.Printf("MEDIA PAR = %.2f\n", mediapar)
    fmt.Printf("MEDIA IMPAR = %.2f\n", mediaimpar)
}
