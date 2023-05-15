package main

import "fmt"

func main() {
	  // Memory -> endress -> value
	  a := 10
	  println(&a)
	  var ponteiro *int = &a // * = endress -> que esta na memoria o ponteiro e um enderecamento na memoria
	  *ponteiro = 20
	  fmt.Println(*ponteiro) //* = mostra tambem o valor que esta na memoria
	  // quando passa um valor de um lugar pro outro, voce esta passando uma copia da memoria
	  
	  // se coloar o ponteiro * nao sera uma copia da memoria sera o valor real e pra passa essa referencia da memoria, bastar colocar &

	  // quando nao usar ponteiro = quando voce so que passa uma copia dos dados utilizar aqueles dados e pronto

	  // quando utilizar o ponteiro = se quiser tornar o valor  mutavel basta utilizar o ponteiro
	  // para ter essa mudanca se precisa utilizar o mesmo valor em varios lugares utilizar o ponteiro

}