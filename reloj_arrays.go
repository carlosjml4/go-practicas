package main

import (
	"fmt"
	"strings"
)

type nrlj struct{ n [][]string }

func IniA(n int, a [][]string) {
	s := "█"

	r := " "
	for i := range a {
		a[i] = make([]string, 5)
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			a[i][j] = r
		}
	}
	switch n {
	case 0:
		a0 := a
		a0[0][1] = s
		a0[0][2] = s
		a0[0][3] = s
		a0[1][1] = s
		a0[1][3] = s
		a0[2][1] = s
		a0[2][3] = s
		a0[3][1] = s
		a0[3][3] = s
		a0[4][1] = s
		a0[4][2] = s
		a0[4][3] = s
	case 1:
		a1 := a
		a1[0][1] = s
		a1[0][2] = s
		a1[1][2] = s
		a1[2][2] = s
		a1[3][2] = s
		a1[4][1] = s
		a1[4][2] = s
		a1[4][3] = s
	case 2:
		a2 := a
		a2[0][1] = s
		a2[0][2] = s
		a2[0][3] = s
		a2[1][3] = s
		a2[2][2] = s
		a2[3][1] = s
		a2[4][1] = s
		a2[4][2] = s
		a2[4][3] = s
	case 3:
		a3 := a
		a3[0][1] = s
		a3[0][2] = s
		a3[0][3] = s
		a3[1][3] = s
		a3[2][2] = s
		a3[3][3] = s
		a3[4][1] = s
		a3[4][2] = s
		a3[4][3] = s
	case 4:
		a4 := a
		a4[0][1] = s
		a4[0][3] = s
		a4[1][1] = s
		a4[1][3] = s
		a4[2][1] = s
		a4[2][2] = s
		a4[2][3] = s
		a4[3][3] = s
		a4[4][3] = s
	case 5:
		a5 := a
		a5[0][1] = s
		a5[0][2] = s
		a5[0][3] = s
		a5[1][1] = s
		a5[2][1] = s
		a5[2][2] = s
		a5[2][3] = s
		a5[3][3] = s
		a5[4][1] = s
		a5[4][2] = s
		a5[4][3] = s
	case 6:
		a6 := a
		a6[0][1] = s
		a6[0][2] = s
		a6[0][3] = s
		a6[1][1] = s
		a6[2][1] = s
		a6[2][2] = s
		a6[2][3] = s
		a6[3][1] = s
		a6[3][3] = s
		a6[4][1] = s
		a6[4][2] = s
		a6[4][3] = s
	case 7:
		a7 := a
		a7[0][1] = s
		a7[0][2] = s
		a7[0][3] = s
		a7[1][3] = s
		a7[2][2] = s
		a7[3][2] = s
		a7[4][2] = s
	case 8:
		a8 := a
		a8[0][1] = s
		a8[0][2] = s
		a8[0][3] = s
		a8[1][1] = s
		a8[1][3] = s
		a8[2][1] = s
		a8[2][2] = s
		a8[2][3] = s
		a8[3][1] = s
		a8[3][3] = s
		a8[4][1] = s
		a8[4][2] = s
		a8[4][3] = s
	case 9:
		a9 := a
		a9[0][1] = s
		a9[0][2] = s
		a9[0][3] = s
		a9[1][1] = s
		a9[1][3] = s
		a9[2][1] = s
		a9[2][2] = s
		a9[2][3] = s
		a9[3][3] = s
		a9[4][1] = s
		a9[4][2] = s
		a9[4][3] = s
	case 11: //:
		a11 := a
		a11[1][2] = s
		a11[3][2] = s
	case 12: // /
		a12 := a
		a12[0][4] = s
		a12[1][3] = s
		a12[2][2] = s
		a12[3][1] = s
		a12[4][0] = s
	}
}

func ImpA(a [][]string) {
	for _, v := range a {
		fmt.Println(strings.Join(v, ""))
	}
}

func main() {
	//inicializa array 2D

	//anomb := []string{"a0", "a1", "a2", "a3", "a3", "a4", "a5", "a6", "a7", "a8", "a9", "a0"}
	//tiempo := []nrlj{}
	//var elrlj nrlj
	a0 := make([][]string, 5)
	a1 := make([][]string, 5)
	a2 := make([][]string, 5)
	a3 := make([][]string, 5)
	a4 := make([][]string, 5)
	a5 := make([][]string, 5)
	a6 := make([][]string, 5)
	a7 := make([][]string, 5)
	a8 := make([][]string, 5)
	a9 := make([][]string, 5)
	a11 := make([][]string, 5)
	a12 := make([][]string, 5)

	IniA(0, a0) //numeros ...
	IniA(1, a1)
	IniA(2, a2)
	IniA(3, a3)
	IniA(4, a4)
	IniA(5, a5)
	IniA(6, a6)
	IniA(7, a7)
	IniA(8, a8)
	IniA(9, a9)
	IniA(11, a11) // :puntos
	IniA(12, a12) // /diagonal

	fmt.Println("------------------")

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print(a1[i][j])

		}
		for j := 0; j < 5; j++ {
			fmt.Print(a11[i][j])

		}
		for j := 0; j < 5; j++ {
			fmt.Print(a2[i][j])

		}
		for j := 0; j < 5; j++ {
			fmt.Print(a12[i][j])

		}
		for j := 0; j < 5; j++ {
			fmt.Print(a4[i][j])

		}
		for j := 0; j < 5; j++ {
			fmt.Print(a5[i][j])

		}
		fmt.Println()
	}

}

//ascii 9608 █
//for _, v := range a0 {
//		fmt.Println(strings.Join(v, ""))
//	}
