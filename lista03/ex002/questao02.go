package main

import "fmt"

func main() {

	var n, x int

	fmt.Println("Entrada")

	for {
		fmt.Scanln(&n)
		if n >= 1 && n <= 1000 {
			break
		}
		fmt.Println("Digito Inválido")
	}

	numeros := make([]int, n, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&numeros[i])
	}
	fmt.Scanln(&x)
	fmt.Println("Saída")

	soma := 0
	for _, v := range numeros {
		if v >= x {
			soma++
		}

	}
	fmt.Println(soma)
}

