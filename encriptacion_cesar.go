package main

import (
	"fmt"
)

func main() {
	/*var text, textc string
	var numc int
	fmt.Println("ingrese el texto")
	fmt.Scan(&text)
	fmt.Println("Ingrese numero")
	fmt.Scan(&numc)*/
	var textc string
	text := "Hola como has estado"
	numc := 3

	for i := 0; i < len(text); i++ {
		if text[i] != 32 {
			textc += string(int(text[i]) + numc)
		} else {
			textc += string(text[i])
		}
	}
	fmt.Println(textc)
}
