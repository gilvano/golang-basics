package main

import "fmt"

func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return // retorno limpo
}

func main() {
	r1, r2 := trocar(1, 2)
	fmt.Printf("r1 %d - r2 %d", r1, r2)
}
