package main

import (
	"fmt"
	"strings"
)

//fb =indice fila blanca, cn=indice columna negra ....[fb][cb])
var ctrl int //control de llamadas max =1

func Mov(f1, f2 string, fb, cb, fn, cn int, atbl [][]string) {
	var s, col, f string

	switch {
	case f1 == "P1" || f1 == "P2": //peones
		if f1 == "P1" && (fn == fb-1 && (cn == cb-1 || cn == cb+1)) {
			s, f = "si", "P"
		} else if f1 == "P2" && (fb == fn+1 && (cb == cn-1 || cb == cn+1)) {
			s, f = "si", "P"
		} else {
			s, f = "no", "P"
		}

		col = DefCol(f1, "P1")
		fmt.Println(f+col, s, " puede comer")

	case f1 == "A1" || f1 == "A2": //alfiles
		i := 0
		for i < 7 {
			if (fn == fb-i && (cn == cb-i || cn == cb+i)) || fn == fb+i && (cn == cb-1 || cn == cb+i) {
				s, f = "si", "A"
				i = 7
			} else {
				i++
				s, f = "no", "A"
			}
		}

		col = DefCol(f1, "A1")
		fmt.Println(f+col, s, " puede comer")

	case f1 == "T1" || f1 == "T2": //torres
		i := 0
		for i < 7 {
			if (fn == fb && (cn == cb-i || cn == cb+i)) ||
				cn == cb && (fn == fb-i || fn == fb+i) {
				s, f = "si", "T"
				i = 7
			} else {
				i++
				s, f = "no", "T"
			}
		}

		col = DefCol(f1, "T1")
		fmt.Println(f+col, s, " puede comer")

	case f1 == "C1" || f1 == "C2": //caballos

		if (fn == fb-2 && (cn == cb-1 || cn == cb+1)) ||
			fn == fb-1 && (cn == cb-2 || cn == cb+2) ||
			fn == fb+1 && (cn == cb-2 || cn == cb+2) ||
			fn == fb+2 && (cn == cb-1 || cn == cb+1) {
			s, f = "si", "C"
		} else {
			s, f = "no", "C"
		}

		col = DefCol(f1, "C1")
		fmt.Println(f+col, s, " puede comer")

	case f1 == "RNA1" || f1 == "RNA": //reinas
		i := 0
		for i < 7 {
			if fn == fb && (cn == cb-i || cn == cb+i) ||
				cn == cb && (fn == fb-i || fn == fb+i) ||
				fn == fb-i && (cn == cb-i || cn == cb+i) ||
				fn == fb+i && (cn == cb-1 || cn == cb+i) {
				s, f = "si", "RNA"
				i = 7
			} else {
				i++
				s, f = "no", "RNA"
			}
		}

		col = DefCol(f1, "RNA1")
		fmt.Println(f+col, s, " puede comer")

	case f1 == "REY1" || f1 == "REY2": //rey
		if fn == fb && (cn == cb-1 || cn == cb+1) ||
			cn == cb && (fn == fb-1 || fn == fb+1) ||
			fn == fb-1 && (cn == cb-1 || cn == cb+1) ||
			fn == fb+1 && (cn == cb-1 || cn == cb+1) {
			s, f = "si", "REY"
		} else {
			s, f = "no", "REY"
		}

		col = DefCol(f1, "REY1")
		fmt.Println(f+col, s, " puede comer")

	}
	ctrl++
	aux := f2
	if ctrl < 2 { //cabiamos de orden fichas y mandamos nuevamente
		f2 = f1
		f1 = aux
		Mov(f1, f2, fb, cb, fn, cn, atbl)
	}
}

func DefCol(f1, s string) string {
	if f1 == s {
		return "B"
	} else {
		return "N"
	}
	return ""
}

func ImpAje(atbl [][]string, f1, f2 string, fn, cn, fb, cb int) {
	st := "█"
	st1 := ""
	st2 := ""
	//♛ ♚ ♝ ♞ ♜ ♟
	//♔ ♕ ♖ ♗ ♘ ♙
	r := " "
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if (i+j)%2 == 0 {
				atbl[i][j] = r
			} else {
				atbl[i][j] = st
			}
		}
	}

	if f1 == "P1" {
		st1 = "♙"
	} else if f1 == "A1" {
		st1 = "♗"
	} else if f1 == "C1" {
		st1 = "♘"
	} else if f1 == "T1" {
		st1 = "♖"
	} else if f1 == "RNA1" {
		st1 = "♕"
	} else if f1 == "REY1" {
		st1 = "♔"
	}

	if f2 == "P2" {
		st2 = "♟"
	} else if f2 == "A2" {
		st2 = "♝"
	} else if f2 == "T2" {
		st2 = "♜"
	} else if f2 == "C2" {
		st2 = "♞"
	} else if f2 == "RNA2" {
		st2 = "♚"
	} else if f2 == "REY2" {
		st2 = "♛"
	}

	atbl[fb][cb] = st1
	atbl[fn][cn] = st2

	for _, v := range atbl {
		fmt.Println(v)
	}
}

func main() {
	var f1, f2 string
	var fb, cb, fn, cn int
	atbl := make([][]string, 8)
	for i := range atbl {
		atbl[i] = make([]string, 8)
	}

	fmt.Println("Fichas P=peon, C= caballo, A=alfil etc. La primera ficha sera ")
	fmt.Println("la Blanca. En este tablero las Blancas avanzan de abajo -> arriba")
	fmt.Println("Ingrese 1a. ficha, luego sus coordenadas separadas por espacio")
	fmt.Print("Primera ficha :")
	fmt.Scan(&f1)
	f1 = strings.ToUpper(f1) + "1"
	fmt.Print("Coord. primera ficha :")
	fmt.Scan(&fb, &cb)
	fmt.Print("Segunda ficha : ")
	fmt.Scan(&f2)
	f2 = strings.ToUpper(f2) + "2"
	fmt.Print("Coord. Segunda ficha :")
	fmt.Scan(&fn, &cn)
	ImpAje(atbl, f1, f2, fn, cn, fb, cb)
	Mov(f1, f2, fb, cb, fn, cn, atbl)
}
