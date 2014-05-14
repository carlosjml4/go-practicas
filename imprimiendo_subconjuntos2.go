package main

import "fmt"

func frecursiva(numero int, indice int, longitud int, K int, N int) {

	i := 0
	if longitud == K {
		fmt.Println(numero)
	} else {
		for i = indice; i <= N; i++ {
			frecursiva(numero*10+i, i+1, longitud+1, K, N)
		}
	}
}

func main() {
	var i, K, N int
	N = 3
	for j := 1; j <= N; j++ {
		K = j - 1
		for i = 1; i <= N; i++ {
			frecursiva(i, i+1, 1, K, N)
		}
	}
}
