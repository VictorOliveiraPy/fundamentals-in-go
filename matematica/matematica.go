package matematica

// com inicial Maisculo todos enxegar o pacote fora do matematica

func Soma[T int | float64](a, b T) T {
	return a + b
}