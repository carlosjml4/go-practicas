//ordena y pasa el 1er y 2° menor

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func n_repetido(n int) bool {
	f := strconv.Itoa(n)
	num := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	ccount := 0
	c := strings.Count(f, num[0])

	for i := 0; i < 10; i++ {
		c = strings.Count(f, num[i])
		if c > 1 {
			ccount = c
		}
		if ccount > 1 {
			return false
			break
		}
	}
	return true
}

func n_div(n int) bool {
	for i := 1; i < 13; i++ {
		if n%i != 0 {
			return false
			break
		}
	}
	return true
}

func main() {
	
	for i := 12340; i < 89765; i++ {
		if n_repetido(i) && n_div(i) {
			fmt.Println(i)
		}
	}
}
