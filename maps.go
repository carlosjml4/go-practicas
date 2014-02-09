package main

import "fmt"

func main() {
	m := map[int]string{1: "lunes", 2: "martes", 3: "miercoles"}
	uno := m[3]

	fmt.Println(m)
	fmt.Println(uno)
	m[4] = "jueves"

	fmt.Println(m)
	x := 5
	var valor string
	var presente bool

	valor, presente = m[x]
	fmt.Println(valor)
	fmt.Println(presente)
}
