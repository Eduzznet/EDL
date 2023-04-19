package main

import "fmt"

type Ints []int

func (a Ints) Len() int            { return len(a) }
func (a Ints) Troca(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Ints) Menor(i, j int) bool { return a[i] < a[j] }
func (a Ints) Maior(i, j int) bool { return a[i] > a[j] }

func Quick(vetor Ints) Ints {
	_quick(vetor, 0, vetor.Len()-1)
	return vetor
}

func _quick(vetor Ints, esquerda int, direita int) {
	if esquerda >= direita {
		return
	}
	pivot := particao(vetor, esquerda, direita)
	_quick(vetor, esquerda, pivot-1)
	_quick(vetor, pivot+1, direita)
}

func particao(vetor Ints, esquerda int, direita int) int {
	pivot := esquerda
	l := esquerda + 1
	r := direita
	for {
		for ; vetor.Menor(l, pivot); l++ {
		}
		for ; vetor.Maior(r, pivot); r-- {
		}
		if l > r {
			break
		}
		vetor.Troca(l, r)
		l++
		r--
	}
	vetor.Troca(esquerda, l-1)
	return l - 1
}

func main() {
	vetor := []int{9, 3, 4, 1, 10, 2, 5, 7, 8, 0, 6}
	fmt.Printf(" vetor original: %v \n", vetor)
	fmt.Printf(" vetor ordenado com QuickSort: %v \n", Quick(vetor))
}
