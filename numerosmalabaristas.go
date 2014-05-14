package main

import (
	. "fmt"
	. "math"
)

func malab(n int, m *[]int) ([]int, int, int) {
	if n%2 == 0 {
		n = int(Trunc((Pow(float64(n), (0.5)))))
	} else {
		n = int(Trunc((Pow(float64(n), (1.5)))))
	}
	*m = append(*m, n)
	if n != 1 {
		malab(n, m)
	}
	max := 0
	for _, v := range *m {
		if v > max {
			max = v
		}
	}
	return *m, len(*m), max
}

func main() {
	lj := []int{}
	l, max, n := 0, 0, 0
	Print("Entre numero entero, mayor a 0 : ")
	Scan(&n)
	lj, l, max = malab(n, &lj)
	Println("Secuencia Malabarista: ", lj)
	Println("Longitud de la serie: ", l)
	Println("Mayor valor: ", max)

}
