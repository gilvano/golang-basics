package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}
type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	f1 := ferrari{"F40", false, 0}
	f1.ligarTurbo()

	var f2 esportivo = &ferrari{"F40", false, 0}
	f2.ligarTurbo()

	fmt.Println(f1, f2)
}
