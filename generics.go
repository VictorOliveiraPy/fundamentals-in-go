package main


type Number interface {
	~int | ~float64
	}

func Soma[T Number](m map[strinf]T) T{
	var soma T

    for _, v := range m {
        soma += v
    }
    return soma
    
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
    }
	return false

}


func main() {
	m := map[string]int{"victor": 1000, "Hugo": 1000, "GitHub": 1000}
	m2 := map[string]float64{"victor": 1000.10, "Hugo": 1000.20, "yasmin": 300.0}
	println(soma(m))
	println(soma(m2))

	}