package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//N:=0
	numeros := []int{}
	text := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese numero")
	for text.Scan() {
		if text.Text() == "/" {
			break
		}
		a := text.Text()
		i, _ := strconv.Atoi(a)

		numeros = append(numeros, i)
	}
	fmt.Println(numeros)
}
