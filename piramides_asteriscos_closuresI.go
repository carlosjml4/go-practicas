package main

import "fmt"

func main() {
	var num, esp, lin, limit2 int
	var ImprInterior func()
	limit2 = 2

	fmt.Println("ingrese un numero")
	fmt.Scan(&num)

	if num%2 == 0 {
		lin = num / 2 //num lineas
	} else {
		lin = num/2 + 1
	}
	esp = lin - 1 //arrobas

	//probando clausuras-----------------------------------
	ImprPiramide := func() {
		for i := 0; i < esp; i++ {
			fmt.Print(" ")
		}
		if num%2 == 0 { //num pares
			fmt.Println("**")
		} else { //num impares
			fmt.Println("*")
		}

		ImprInterior = func() {
			if esp > 1 {
				limit := esp
				for i := 0; i < limit-1; i++ {
					fmt.Print(" ")
				}
				fmt.Print("*")

				if num%2 == 0 {
					for i := 0; i < limit2; i++ {
						fmt.Print("@")
					}
					fmt.Println("*")
				} else {
					for i := 0; i < limit2-1; i++ {
						fmt.Print("@")
					}
					fmt.Println("*")
				}
				limit2 += 2
				esp--
				ImprInterior()
			}
		}

		ImprInterior()

		for i := 0; i < num; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	//---------------------------------------------------------
	ImprPiramide()

}
