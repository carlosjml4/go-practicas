package main

import "fmt"

func SubConj(a []string) {
	var atmp = []string{}
	var m func(s string, indice int, len_s int, K int)

	m = func(s string, indice int, len_s int, K int) {
		if len_s == K+1 {
			atmp = append(atmp, s)
		} else {
			for i := indice; i <= len(a); i++ {
				m(s+a[i-1], i+1, len_s+1, K)
			}
		}
	}

	for j := 1; j <= len(a); j++ {
		K := j - 1
		for i := 1; i <= len(a); i++ {
			m(a[i-1], i+1, 1, K)
		}
	}

	fmt.Println(atmp)
}

func main() {
	al := make([]string, 5)
	al = []string{"A", "B", "C", "D", "E"}
	SubConj(al)
}
