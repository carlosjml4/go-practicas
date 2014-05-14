//t=terminos y = exponente +1
//indice = n!/((n-i)!*i!), donde n es el exponente
package main

import (
	. "fmt"
	. "math"
)

func facto(nf int) float64 {
	if nf == 0 {
		return 1
	}
	return float64(nf) * facto(nf-1)
}
func impasc(n int, s string) {
	if s == "t" {
		for k := 0; k <= n; k++ {
			for j := n - k; j > 0; j-- {
				Print(" ")
			}
			for i := 0; i <= k; i++ {
				Print(Trunc(facto(k)/(facto(k-i)*facto(i))), " ")
			}
			Println()
		}
	} else {
		for i := 0; i <= n; i++ {
			Print(Trunc(facto(n)/(facto(n-i)*facto(i))), " ")
		}
	}
}

func main() {
	n := 0
	s := ""
	Print("Introduzca fila: ")
	Scan(&n)
	Print("Quieres imprimir la  fila o toda la piramide (f o t): ")
	Scan(&s)
	impasc((n - 1), s)
	Println()
}
