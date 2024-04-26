Questão 05

package main

import "fmt"


func maxsegmento(num []int) int {
	if len(num) == 0 {
		return 0
	}

	maxComp := 0
	comp := 0

	
	for i := 1; i < len(num); i++ {
		if num[i] > num[i-1] {
			
			comp++
		} else {
			
			if comp > maxComp {
				maxComp = comp
			}
			comp = 0 
		}
	}

	
	if comp > maxComp {
		maxComp = comp
	}

	return maxComp
}

func main() {
	var n int
	fmt.Scan(&n)

	num := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&num[i])
	}

	maxSegmentoCrescente := maxsegmento(num)

	fmt.Printf("O comprimento do segmento crescente maximo e: %d\n", maxSegmentoCrescente)
}
