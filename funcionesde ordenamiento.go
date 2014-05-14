package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int

	fmt.Println("ingrese 10 numeros")
	s := []int{}
	for i := 0; i < 10; i++ {
		fmt.Scan(&n)
		s = append(s, n)
	}
	sort.Insts(s)
	fmt.Println(s)
}
