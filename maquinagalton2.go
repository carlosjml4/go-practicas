package main

import (
	. "fmt"
	"math/rand"
	"time"
)

//generador de numeros al azar **copperMan**
func randInt() int {
	rand.Seed(time.Now().UTC().UnixNano())
	nrd := -1 + rand.Intn(3)
	for nrd == 0 {
		nrd = -1 + rand.Intn(3)
	}
	return nrd
}
func galton(a []int) {

}

func main() {
	bal := 0

	a1 := []int{}
	a2 := []int{}
	a3 := []int{}
	a4 := []int{}
	a5 := []int{}
	a6 := []int{}
	a7 := []int{}
	a8 := []int{}
	a9 := []int{}
	a10 := []int{}
	a11 := []int{}
	a12 := []int{}
	a13 := []int{}

	for i := 0; i < 10000; i++ {
		//piramide de 12 pivotes de base
		//la "bola" empieza a caer por la piramide
		//tomando valores al azarcon +1 o -1
		for j := 0; j < 14; j++ {
			bal += randInt()
		}
		//almacena en cada "recipeinte" segun el valor final
		switch bal {
		case -12:
			a1 = append(a1, bal)
		case -10:
			a2 = append(a2, bal)
		case -8:
			a3 = append(a3, bal)
		case -6:
			a4 = append(a4, bal)
		case -4:
			a5 = append(a5, bal)
		case -2:
			a6 = append(a6, bal)
		case 0:
			a7 = append(a7, bal)
		case 2:
			a8 = append(a8, bal)
		case 4:
			a9 = append(a9, bal)
		case 6:
			a10 = append(a10, bal)
		case 8:
			a11 = append(a11, bal)
		case 10:
			a12 = append(a12, bal)
		case 12:
			a13 = append(a13, bal)
		}
		bal = 0
	}
	Println()
	//imprime resultado teniendo en cuenta la cantidad de elementos "bal"
	//que se acumularon en cada recipiente osea el "len"" de cada array"
	Println("Los valores para el histograma son: ")
	Printf("| %d | %d | %d | %d | %d | %d | %d | %d | %d | %d | %d | %d | %d |",
		len(a1), len(a2), len(a3), len(a4), len(a5), len(a6), len(a7), len(a8),
		len(a9), len(a10), len(a11), len(a12), len(a13))
}
