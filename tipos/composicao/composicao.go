package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	abrirTeto()
}

type bmw struct{}

func (b bmw) ligarTurbo() {
	fmt.Println("Turbo...")
}

func (b bmw) fazerBaliza() {
	fmt.Println("Baliza...")
}

func (b bmw) abrirTeto() {
	fmt.Println("Teto...")
}

func main() {
	var b esportivoLuxuoso = bmw{}
	b.ligarTurbo()
	b.fazerBaliza()
	b.abrirTeto()
}
