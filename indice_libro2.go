package main

import (
	. "fmt"
	. "strconv"
)

type ele_ind struct {
	titulo string
	nivel  int
}

//devuelve 10,100,1000.......
func ceros(n int) (dec int) {
	var s string = "1"
	for i := 1; i < n; i++ {
		s = s + "0"
	}
	dec, _ = Atoi(s)
	return
}
//pone puntos en numeros de indice
func pp(st string)s string{
    for _,v:=range st{
    fmt.Print(string(v),".")
}
}

//crea array con numeros de indice
func cind(a []ele_ind, ai *[]int) {
	for _, v := range a {
		*ai = append(*ai, v.nivel)
	}
}

//crea array con arbol de numeros de indice
func ctind(ai []int, ind *[]string) {
	s := ""
	ant := make(map[int]string)
	ante := 0
	i := 1

	for _, v := range ai {
		if ante == 0 {
			ante = v
			s = Itoa(v)
			*ind = append(*ind, s)
			ant[v] = s
		} else if v > ante {
			i = 1
			ante = v
			s = ant[v-1] + Itoa(i)
			*ind = append(*ind, s)
			ant[v] = s
			i++
		} else if v == ante {
			s = ant[v-1] + Itoa(i)
			*ind = append(*ind, s)
			ant[v] = s
			i++
		} else if v < ante {
			i, _ = Atoi(ant[v])
			if v != 1 {
				ante = v
				i = (i % ceros(v)) + 1
				s = ant[v-1] + Itoa(i)
				*ind = append(*ind, s)
				ant[v] = s
				i++
			} else {
				ante = v + 1
				i = i + 1
				s = Itoa(i)
				*ind = append(*ind, s)
				ant[v] = s
				i = 1
			}

		}
	}
}

//crea el numero de espacios segun el nivel
//devolviendo una cadena de espacios
func tab(n int) string {
	esp := " "
	for i := 0; i < n; i++ {
		esp = esp + "  "
	}
	return esp
}

//imprime tabla contenido
func tcont(a []ele_ind, ind []string) {
	for i, v := range a {
		Println(tab(v.nivel), ind[i], ".", v.titulo)
	}
}

func main() {
	ai := []int{}
	ind := []string{}
	cont := []ele_ind{{"Introducción", 1},
		{"Motivación", 2},
		{"Reseña histórica", 2},
		{"Origen", 3},
		{"Trabajos preliminares", 3},
		{"Soluciones actuales", 2},
		{"Objetivo de este libro", 2},
		{"Requisitos", 1},
		{"Hardware", 2},
		{"Software", 2}}
	cind(cont, &ai)
	ctind(ai, &ind)
	tcont(cont, ind)
}
