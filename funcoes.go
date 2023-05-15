package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum_bol(20, 0))

}

func sum(a, b int) int { // se for o mesmo tipo os argumentos
	return a + b
}


func sum_bol(a, b int) (int, bool) { // se for o mesmo tipo os argumentos
	if a + b >= 30 {
		return a + b, true
	}

	return a + b, false
}