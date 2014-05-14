package main

import (
	"fmt"
	//"math"
	"strconv"
)

func EsFeliz(num int) int {
	var sum int
	nums := strconv.Itoa(num)
	lnum := len(nums)

	switch lnum {
	case 1:
		{
			if num == 1 {
				return 1
			} else {
				sum = (num * num)
				return sum
			}
		}
	case 2:
		{
			dec := num / 10
			sum = (dec*dec + EsFeliz(num%10))
			return sum
		}
	case 3:
		{
			cen := num / 100
			sum = (cen*cen + EsFeliz(num%100))
			return sum
		}
	case 4:
		{
			mil := num / 1000
			sum = (mil*mil + EsFeliz(num%1000))
			return sum
		}
	}
	return 0
}

func main() {
	var num, nnum int
	b := false
	fmt.Println("Ingrese un numero entero entre 1 y 9999 ")
	fmt.Scan(&num)
	if num > 9999 {
		fmt.Println("fuera de rango")
		num = 0
	}
	fmt.Print(num)
	nnum = EsFeliz(num)

	for i := 0; i < 20; i++ {
		if nnum == 1 {
			b = true
			break
		} else {
			nnum = EsFeliz(nnum)
		}
	}

	if b == true {
		fmt.Println(", Es un numero Feliz ;-))")
	} else {
		fmt.Println(", No se ha podido determinar ;-))")
	}
	main()
}
