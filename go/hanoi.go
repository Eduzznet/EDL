package main

import "fmt"

func Hanoi(disco, origem, destino, aux int) {
	if disco <= 0 {
		return
	}
	Hanoi(disco-1, origem, aux, destino)
	fmt.Println(" Mover disco ", disco, " de [", origem, "] para  [", destino, "]")
	Hanoi(disco-1, aux, destino, origem)
}

func main() {
	var disco int = 4
	Hanoi(disco, 1, 3, 2)
}
