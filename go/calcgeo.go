package main

import (
	"fmt"
	"math"
)

type geometria interface {
	area() float64
	perim() float64
}

type rect struct {
	largura, altura float64
}
type circulo struct {
	raio float64
}

func (r rect) area() float64 {
	return r.largura * r.altura
}
func (r rect) perim() float64 {
	return 2*r.largura + 2*r.altura
}

func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}
func (c circulo) perim() float64 {
	return 2 * math.Pi * c.raio
}

func medidas(g geometria) {
	fmt.Println("Dimensões: ", g)
	fmt.Println("Area: ", g.area())
	fmt.Println("Perimetro: ", g.perim())
}

func main() {
	r := rect{largura: 3, altura: 4}
	c := circulo{raio: 5}

	fmt.Println("Medidas de um Retangulo:")
	medidas(r)
	fmt.Println("Medidas de um Circulo:")
	medidas(c)
}
