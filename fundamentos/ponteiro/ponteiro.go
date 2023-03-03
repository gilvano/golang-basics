package main

import "fmt"

func main() {
	i := 1

	// Go não tem aritmética de ponteiros
	var p *int = nil

	p = &i // pega o endereço de i
	*p++   // desreferencia p e incrementa
	fmt.Println(*p, i)
}
