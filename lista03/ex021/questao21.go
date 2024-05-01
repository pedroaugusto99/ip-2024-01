package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Println("Entrada")
	fmt.Scanln(&n)
	s := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&s[i])
	}
	fmt.Println("Saída")

	sort.Ints(s)

	for _, v := range s {
		if v%2 == 0 {
			fmt.Printf("%v\n", v)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(s)))

	for _, v := range s {
		if v%2 != 0 {
			fmt.Printf("%v\n", v)
		}
	}
}
