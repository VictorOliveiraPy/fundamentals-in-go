package main

import "fmt"


type Conta struct {
	saldo float64

}

func NewConta(saldo float64) *Conta {
	return &Conta{saldo: saldo}
}

func (c *Conta) simular(valor float64) float64 {
	c.saldo += valor
	return c.saldo
}

func main() {
	c := Conta{saldo: 200}
    c.simular(100)
    fmt.Println(c.saldo)


}