package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[12345] = "Maria"
	aprovados[12333] = "Pedro"
	aprovados[12222] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12222])

	delete(aprovados, 12333)

	fmt.Println(aprovados)
}
