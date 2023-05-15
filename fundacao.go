package main

import (
	"fmt"
)

type ID int

var (
	f ID = 1
)


func main(){
	var  myArray [3] int
	myArray[0] = 10
	myArray[1] = 22
	myArray[2] = 30


	fmt.Println(myArray[len(myArray) -1]) // pegar o ultimo elemento

	for i, v := range myArray {
		fmt.Printf("O valor de %d e %d", i, v)
		
	}

	s := []int{1, 2, 3, 4}
	fmt.Printf("len=%d  cap=%d %v", len(s), cap(s), s)


}

// pegar o tipo de dados %T, pegar o valor da var %v




// Percorrer array