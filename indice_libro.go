package main

import . "fmt"

type ele_ind struct {
	titulo string
	nivel  int
}
//definimos funcion como tipo variable
//para colocar espacios de tabulacion 
type tb func(int) string

//crea el numero de espacios segun el nivel
//devolviendo una cadena de espacios
func tab(n int) string {
	esp := " "
	for i := 0; i < n; i++ {
		esp = esp + "  "
	}
	return esp
}

//imprime el arbol
func indice(a []ele_ind, f tb) {
	i_t := 0
    s_t:=""
	for _, v := range a {
		if v.nivel == 1 {
			i_t = i_t + 1
		}
		Println(tab(v.nivel), i_t, ".", v.titulo)
	}
}

func main() {
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
	indice(cont, tab)
}
