package main

import (
	"fmt"
)

//var message string = "hello world"
func main() {
	//:=で
	message := "hello world"
	fmt.Println(message)

	//constは定数になる
	const Hello string = "hello"
        Hello = "bye" // cannot assign to Hello
}
