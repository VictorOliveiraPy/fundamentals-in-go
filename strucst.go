package main

import "fmt"

type Endereco struct {
	Longradouro string
	Numero int
	Cidade string
	Estado string
}

type Client struct {
	Nome string
	Idade int
	Ativo bool
	Endereco
}

func(c Client) Desativar(){
	c.Ativo = false
	fmt.Printf("O client: %s foi desativado\n", c.Nome)
}


func main() {
	victor := Client{
		Nome: "Victor",
        Idade: 20,
        Ativo: true,
	}
	victor.Cidade = "Arcoverde"
	fmt.Println(victor)
	victor.Desativar()

}