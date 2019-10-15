package main

import (
  "fmt"

)


func hello() {
  fmt.Println("hello")
}

func sum(i, j int) {
   fmt.Println(i + j)
}

func sum2(i, j int) int {

   return i + j

}

func swap(a, b int) (int, int) {
   return a, b

}

func main() {
    hello()

    sum(1, 2)

    n := sum2(3, 4)
    fmt.Println(n)
    
    x, y := 1, 2
    x, y = swap(x, y)
    fmt.Println(x,y)

  //  x = swap(x,y) コンパイルエラー

   x, _ = swap(x, y)
   fmt.Println(x)


   swap(x,y)
}
