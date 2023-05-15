package main

import "fmt"

// o tipo da interface vazia aceita qualquer coisar

func main() {
	var minhaVar interface{} = "Victor Hugo"
	println(minhaVar.(string))
	res, ok = minhaVar.(int)
	fmt.Printf("O valor de res e %v\n e o resultado de ok e %v", res, ok)
}

