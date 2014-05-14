package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ABis(a int) bool {
	if a%4 == 0 && (a%100 != 0 || a%400 == 0) {
		return true
	} else {
		return false
	}
}

func MesDias(m, a int) int {
	switch {
	case m == 1 || m == 3 || m == 5 || m == 7 || m == 8 || m == 10 || m == 12:
		{
			return 31
		}
	case m == 4 || m == 6 || m == 9 || m == 11:
		{
			return 30
		}
	case m == 2:
		{
			if ABis(a) {
				return 29
			} else {
				return 28
			}
		}
	}
	return 0
}

func main() {
	var fech, fechan string

	fmt.Println("Ingrese una fecha (dia,mes,año)")
	fmt.Scan(&fech)
	nfech := strings.Split(fech, ",")
	d, _ := strconv.Atoi(nfech[0])
	m, _ := strconv.Atoi(nfech[1])
	a, _ := strconv.Atoi(nfech[2])

	if d == 1 && m > 1 { 
		fechan = strconv.Itoa(MesDias((m-1), a)) + "," +
			strconv.Itoa(m-1) + "," + strconv.Itoa(a)
	} else if d == 1 && m == 1 { 
		fechan = strconv.Itoa(MesDias((12), a)) + "," +
			strconv.Itoa(12) + "," + strconv.Itoa(a-1)
	} else {
		fechan = strconv.Itoa(d-1) + "," + strconv.Itoa(m) + "," +
			strconv.Itoa(a)
	}
	fmt.Println("la fecha anterior es: ", fechan)
	main()
}
