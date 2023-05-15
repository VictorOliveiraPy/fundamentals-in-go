package main


// or = || 
// ad = &&

func main() {
	a := 1
	b := 2
	c := 3

	switch {
		case a > b:
            println("a > b")
        case a < b:
            println("a < b")
        case a == b:
            println("a == b")
        case a!= b:
            println("a!= b")
		default:
			println("Default")
	}


	if a >b && c >a {
		println("a > b && c > a")
    } else if a >b && c <a {

	}

	if a == b{
		println(a)
	} else {
		println(b)
	}

}