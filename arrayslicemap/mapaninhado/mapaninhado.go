package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15654.98,
			"Guga Pereira":   12332.09,
		},
		"J": {
			"José João": 12333.32,
		},
		"P": {
			"Pedro Junior": 12000.32,
		},
	}

	for letra, funcs := range funcsPorLetra {
		fmt.Printf("Letra (%s).........\n", letra)
		for fun, sal := range funcs {
			fmt.Printf("%s - %f.2\n", fun, sal)
		}
	}
}
