package main

import(
   "fmt"
)


func main() {
   n:=0
   for {
     n++

     if n > 10 {
	 break
     }

     if n%2 == 0 {
	  continue
     }
     fmt.Println(n)
   }

}
