package main

import "fmt"


func main(){
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14))

}

func sum(numbers ...int) int { // quando nao sabe o numero de quantidades que será passado utilizar ...
	sum := 0
    for _, num := range numbers {
        sum += num
    }
    return sum
}
