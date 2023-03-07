package main

import "fmt"

func main() {
	funcESalarios := map[string]float64{
		"José João":      11325.54,
		"Gabriela Silve": 13221.34,
		"Pedro Junior":   1200.00,
	}

	funcESalarios["Rafael Filho"] = 1350.00

	fmt.Println(funcESalarios)

	delete(funcESalarios, "Inexistente")

	for nome, salario := range funcESalarios {
		fmt.Printf("%s : %f.2\n", nome, salario)
	}
}
