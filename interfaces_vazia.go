package main

import "fmt"

// o tipo da interface vazia aceita qualquer coisar

func main() {
	var x interface{} = 10
	var y interface{} = "Hello, world!"

	showType(x)
	showType(y)


}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel e %T e o valor e %v", t, t)
}