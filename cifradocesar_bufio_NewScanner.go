package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var textc string
	var numc int
	bio := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese el texto")
	bio.Scan()
	text := bio.Text()

	fmt.Println("Ingrese numero")
	fmt.Scan(&numc)

	for i := 0; i < len(text); i++ {
		if text[i] != 32 {
			textc += string(int(text[i]) + numc)
		} else {
			textc += string(text[i])
		}
	}
	fmt.Println(textc)
}
