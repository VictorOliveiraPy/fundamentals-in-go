package main

import "fmt"

/// closures function anonymas


func main() {
	total := func ()int  {
		return sum(1, 2, 3, 4, 534, 32) * 2
		
	}
	fmt.Println(total)

}


func sum(numbers ...int) int { // quando nao sabe o numero de quantidades que será passado utilizar ...
	sum := 0
    for _, num := range numbers {
        sum += num
    }
    return sum
}
