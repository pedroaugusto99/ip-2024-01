package main

import (
  "fmt"
  "sort"
  ) 

func main(){
  var q1, q2 int
  
  fmt.Scan(&q1)
  fmt.Scan(&q2)
  
  V1 := make([]int, q1)
  V2 := make([]int, q2)
  
  for i:=0;i < q1;i++{
    fmt.Scan(&V1[i])
  }
  
  for i := range V2{
    fmt.Scan(&V2[i])
  }
  
  Vr := append(V1, V2... )
  
  sort.Ints(Vr)
  
  for _, v:= range Vr{
  fmt.Println(v)
  }
}
