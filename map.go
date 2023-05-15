package main

import "fmt"

func main() {
	salarios  := map[string]int{"victor": 8200, "suicide":200, "Maria": 3000}
	fmt.Println(salarios["victor"])
	delete(salarios, "suicide")
	salarios["suicide"] = 200

	for nome, salario := range salarios {
		fmt.Println(nome, salario)
			}

	for _, salario := range salarios {
		fmt.Println(salario)
			}
		
}
