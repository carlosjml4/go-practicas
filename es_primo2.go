package main

import (
	. "fmt"
	. "math"
)

type primInt func(int) int // definimos un función como un tipo de variable

func esprim(n int) int {
	for i := 2; i <= int(Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return 0
		}
	}
	return n
}

func rpta(n int, f primInt) (s string) {
	if f(n) > 0 {
		s = "si"
	} else {
		s = "no"
	}
	return s
}

func main() {
	n := 0
	Println(n)
	Println("Digite un numero entero mayor a 1: ")
	Scan(&n)
	Println("el numero", n, rpta(n, esprim), "es primo")

}
