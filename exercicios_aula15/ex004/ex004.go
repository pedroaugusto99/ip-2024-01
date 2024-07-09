package main

import (
	"fmt"
)

func decimalParaBinario(decimal int) string {
	if decimal == 0 {
		return "0"
	}
	var binario string
	for decimal > 0 {
		binario = string(decimal%2+'0') + binario
		decimal /= 2
	}
	return binario
}

func main() {
 var numeroDecimal int
	fmt.Scanln(&numeroDecimal
	fmt.Println(decimalParaBinario(numeroDecimal))
}
