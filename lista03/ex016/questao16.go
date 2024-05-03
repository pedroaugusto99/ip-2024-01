package main

import (
  "fmt"
  "sort"
)

func main() {
  var N, K int
  fmt.Scan(&N, &K)

  horarios := make([]int, N)
  for i := 0; i < N; i++ {
    fmt.Scan(&horarios[i])
  }
  
  indices := make(map[int]int)
  for i, tempo := range horarios {
    indices[tempo] = i + 1
  }

  
  presentes := 0
  for _, tempo := range horarios {
    if tempo <= 0 {
      presentes++
    }
  }

  
  if presentes <= K {
    fmt.Println("SIM")
  } else {
    fmt.Println("NÃO")
    
    sort.Ints(horarios)
    //sort.Slice(horarios, func(i, j int) bool {
      return horarios[j] > horarios[i]
    })

    for _, tempo := range horarios {
      if tempo <= 0 {
        fmt.Printf("%d\n", indices[tempo])
      }
      }
    }
    fmt.Println()
  }

