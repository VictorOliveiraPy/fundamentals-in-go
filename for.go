package main


// No Go tem apenas o For 

func main() {
	for i:= 0; i < 10; i++ {
		println(i)

}

	numeros := []string{"um", "dois", "tres"}
	for k, v := range numeros {
		println(k, v)
    }

	i:= 0
	for i < 10 {
        println(i)
        i++
    }

	for {
		println("Hello, world") //looping infinite loop
	}


}

// 819843-53372