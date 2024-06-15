package main
import "fmt"

func main() {
    var x int
    var n int
    fmt.Scanln(&x)
     fmt.Scanln(&n)
    
    
    res := pot(x,n)
    fmt.Println(res)
    
}

func pot(x int, n int) int {
    if n <= 0 {
        return 1
    }
    
    return x * pot(x, n - 1)
}
