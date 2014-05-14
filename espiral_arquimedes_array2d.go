package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//imprime array normal
func a2d_normal(a2 [][]int, f, c int) {
	for i := 0; i < f; i++ {
		for j := 0; j < c; j++ {
			fmt.Println("(", i, ",", j, ")=", a2[i][j])
		}
	}
}

func a2d_arriba(a2 [][]int, f, c int) {
	for i := 0; i < f; i++ {
		for j := 0; j < c; j++ {
			fmt.Println("(", i, ",", j, ")=", a2[i][j])
		}
	}
}

func main() {
	var filas, columnas int
	fmt.Print("Entre un numero para filas y columnas: ")
	fmt.Scan(&filas)
	columnas = filas

	//inicializa array 2D
	a2 := make([][]int, filas)
	for i := range a2 {
		a2[i] = make([]int, columnas)
	}
	//llenamos array
	text := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese los numeros")

	for i := 0; i < filas; i++ {
		for j := 0; j < columnas; j++ {
			text.Scan()
			a := text.Text()
			vlr, _ := strconv.Atoi(a)
			a2[i][j] = vlr
		}
	}

	a2d_normal(a2, filas, columnas)

}
