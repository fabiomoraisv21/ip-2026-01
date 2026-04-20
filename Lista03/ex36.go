package main

import (
	"fmt"
	"strconv"
)

var nBaseDez, quociente int
var rStr, hexa string

func main() {
	fmt.Println("Digite o numero na base 10:")
	fmt.Scan(&nBaseDez)
	lista := []string{}
	hexa = ""
	for {
		quociente = nBaseDez / 16
		res := nBaseDez % 16
		if res >= 0 && res <= 9 {
			rStr = strconv.Itoa(res)
		} else {
			if res == 10 {
				rStr = "A"
			} else if res == 11 {
				rStr = "B"
			} else if res == 12 {
				rStr = "C"
			} else if res == 13 {
				rStr = "D"
			} else if res == 14 {
				rStr = "E"
			} else if res == 15 {
				rStr = "F"
			}
		}
		lista = append(lista, rStr)
		nBaseDez = quociente
		if quociente == 0 {
			break
		}
	}
	for i := len(lista) - 1; i >= 0; i-- {
		hexa += lista[i]
	}
	fmt.Printf("O número em hexadecimal é %v\n", hexa)
}
