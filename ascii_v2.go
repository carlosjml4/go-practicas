//Codigos ascii
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var (
		s     string
		n, nn int
	)

	fmt.Println("Que desea?")
	fmt.Println("1-Numeros")
	fmt.Println("2-Caracteres")
	if _, err := fmt.Scan(&n); err == nil && n <= 2 && n >= 1 {
		switch n {
		case 1:
			fmt.Print("Digite numero =")
			fmt.Scan(&nn)
			fmt.Println("El caracter es =", string(nn))
		case 2:
			fmt.Print("Digite caracter =")
			fmt.Scan(&s)
			r, _ := utf8.DecodeRuneInString(s)
			fmt.Println("El codigo es =", r)
		}
	} else {
		fmt.Println("Debes colocar un numero 1 o 2")
	}
}
