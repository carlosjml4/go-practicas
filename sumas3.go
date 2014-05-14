package main

import "fmt"

func SubConj(as []int, t int) {
	var atmp = []string{}
	var m func(e int, s string, vl, len_s, K int)

	//funcion recursiva que pasa por todos los posibles sumandos
	m = func(e int, s string, vl, len_s, K int) {
		b := false
		if len_s == K+1 {
			//revisa si la cadena al revez (s1)  o si la cadena
			//s clculada ya esxiste en atmp
			arr := []byte(s)
			s1 := string(SRev(arr))
			for _, v := range atmp {
				if v == s1 || v == s {
					b = true
				}
			}
			//--------------------------------------
			if e == t && b == false {
				atmp = append(atmp, s) //si las cadenas no estan las adiciona
			}
		} else {
			for i := vl; i <= len(as); i++ {
				m(e+(as[i-1]), s+"+"+fmt.Sprint(as[i-1]), i+1, len_s+1, K)
			}
		}
	}

	for j := 1; j <= len(as); j++ {
		K := j - 1
		for i := 1; i <= len(as); i++ {
			m(as[i-1], fmt.Sprint(as[i-1]), i+1, 1, K)
		}
	}

	fmt.Println("Para el valor: ", t, atmp)
}

//funcion para volterar la cadena
func SRev(arr []byte) (s string) {
	for i := len(arr) - 1; i >= 0; i-- {
		s += string(arr[i])
	}
	return
}

func main() {
	a_tot := []int{12, 11, 8}
	a_sumd := []int{1, 3, 9, 6, 7, 5}
	for i := 0; i < len(a_tot); i++ {
		SubConj(a_sumd, a_tot[i])
	}
}
