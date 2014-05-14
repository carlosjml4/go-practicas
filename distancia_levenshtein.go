//distancia de caracteres
package main

import (
	"container/list"
	"fmt"
)

var (
	list1 = list.New()
	list2 = list.New()
)

func list_str(s string) {
	for _, v := range s {
		//fmt.Println(i, v)
		list1.PushBack(string(v))
	}
}

func main() {
	distancia := 0
	dif := 0
	it := 0
	st1 := "cazador"
	st2 := "cazar"
	lst1 := len(st1)
	lst2 := len(st2)

	if lst1 < lst2 {
		it = lst1
		dif = lst2 - lst1
	} else {
		it = lst2
		dif = lst1 - lst2
	}

	for i := 0; i < it; i++ {
		if st1[i] != st2[i] {
			distancia++
		}
	}
	fmt.Println(distancia + dif)
	/*fmt.Println("digite 1.a palabra: ")
	fmt.Scan(&st1)
	fmt.Println("digite 2.a palabra: ")
	fmt.Scan(&st2)*/

	/*fmt.Println(lst1, lst2)
	for i, v := range st1 {
		fmt.Println(i, string(v))
		list1.PushBack(string(v))
	}
	list1.PushBack("h")
	fmt.Println(list1.Front().Value)*/
	//list_str(st1)
	//list_str(st2)
	//list2.PushBack(st2[0])

}
