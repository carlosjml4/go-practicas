package main

import (
	"container/list"
	"fmt"
	"math"
)

func es_primo(n int64) bool {
	var i int64
	nr := math.Sqrt(float64(n))
	if n > 1 {
		for i = 2; i <= int64(nr); i++ {
			if n%i == 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	var num, i int64
	var x list.List

	i = 2
	fmt.Println("Digite un numero limite: ")
	fmt.Scan(&num)
	for i < (num + 1) {
		if es_primo(i) {
			x.PushBack(i)
		}
		i++
	}
	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
