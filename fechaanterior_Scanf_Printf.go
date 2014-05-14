package main

import (
	"fmt"
)

func ABis(a int) bool {
	if a%4 == 0 && (a%100 != 0 || a%400 == 0) {
		return true
	}
	return false
}

func MesDias(m, a int) int {
	switch {
	case m == 2:
		{
			if ABis(a) {
				return 29
			} else {
				return 28
			}
		}
	case m%2 != 0 && m < 8 || m%2 == 0 && m > 7:
		{
			return 31
		}
	default:
		{
			return 30
		}
	}
}

func main() {
	var d, m, a int
	var fechan string

	fmt.Println("Ingrese una fecha (dia,mes,año)")
	_, e := fmt.Scanf("%d,%d,%d", &d, &m, &a) //,%2d,%4d", &d, &m, &a)

	if e != nil || (d > 31 || d < 1) || (m > 12 || m < 1) {
		fmt.Println("Valores o formato de entrada errados")
		return
	}

	if d == 1 && m > 1 {
		fechan = fmt.Sprintf("%d,%d,%d", MesDias((m-1), a), (m - 1), a)
	} else if d == 1 && m == 1 {
		fechan = fmt.Sprintf("%d,%d,%d", MesDias(12, a), 12, (a - 1))
	} else {
		fechan = fmt.Sprintf("%d,%d,%d", (d - 1), m, a)
	}
	fmt.Println("la fecha anterior es: ", fechan)
	main()
}
