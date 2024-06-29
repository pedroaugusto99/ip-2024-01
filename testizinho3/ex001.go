package main

import "fmt"

func main() {
	fmt.Scanln(frase)
	
	fmt.Println(Inverter(frase))
}

func Inverter(frase string) (saida string) {
	for _ , v := range frase {
		saida = string(v) + saida
	}
	return
}
