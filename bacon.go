package main

import (
	"fmt"
	"strings"
)

func code(s1, s2 string) {
	//los ascii, para espacio, A y B son 32, 65,66 respectivamente
	var scode, scode2 string
	var isp int //para controlar iteracion con los espacios

	dicbacon := map[string]string{"a": "AAAAA", "b": "AAAAB", "c": "AAABA",
		"d": "AAABB", "e": "AABAA", "f": "AABAB", "g": "AABBA", "h": "AABBB",
		"i": "ABAAA", "j": "ABAAA", "k": "ABAAB", "l": "ABABA", "m": "ABABB",
		"n": "ABBAA", "o": "ABBAB", "p": "ABBBA", "q": "ABBBB", "r": "BAAAA",
		"s": "BAAAB", "t": "BAABA", "u": "BAABB", "v": "BAABB", "W": "BABAA",
		"x": "BABAB", "y": "BABBA", "z": "BABBB"}

	//recorre s1 comparando con el dic y crea el codigo
	for _, v := range s1 {
		scode += dicbacon[string(v)]
	}

	if len(s2) >= len(s2) { //para verificar tamaño texto portador
		//toma numero de espacios y los suma a la longitud
		//del texto a codificar
		for i := 0; i < (len(scode) + strings.Count(s2, " ")); i++ {
			if s2[i] != 32 && scode[isp] == 66 {
				scode2 = scode2 + strings.ToUpper(string(s2[i]))
				isp++
			} else if s2[i] != 32 && scode[isp] == 65 {
				scode2 = scode2 + strings.ToLower(string(s2[i]))
				isp++
			} else if s2[i] == 32 {
				scode2 = scode2 + string(s2[i])
			}
		}
		fmt.Println(scode)
		//escribe los caracteres restantes del texto portador
		for i := (len(scode) + strings.Count(s2, " ")); i < len(s2); i++ {
			scode2 += string(s2[i])
		}
	}
	fmt.Println(scode2)
}

func main() {

	s1 := "python"
	s2 := "texto para portar un mensaje escondido"
	code(s1, s2)
}
