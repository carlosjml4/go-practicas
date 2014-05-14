package main

import "fmt"

func Parimpar(n int) {
	ia := []int{} //impares
	pa := []int{} //pares

	for i := 1; i < n; i++ {
		if i%2 == 0 {
			pa = append(pa, i)
		} else {
			ia = append(ia, i)
		}
	}
	fmt.Println("listado de pares :", pa, len(pa), " numeros")
	fmt.Println("listado de impares :", ia, len(ia), " numeros")
}
func main() {
	n := 0
	fmt.Print("Ingrese numero limite: ")
	fmt.Scan(&n)
	Parimpar(n)
}
