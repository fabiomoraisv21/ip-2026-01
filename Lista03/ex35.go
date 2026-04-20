package main

import (
	"fmt"
	"strconv"
)

var nBase10, q int
var r, binario string

func main() {
	fmt.Println("Digite o numero na base 10:")
	fmt.Scan(&nBase10)
	lista := []string{}
	binario = ""
	for {
		q = nBase10 / 2
		res := nBase10 % 2
		r = strconv.Itoa(res)
		lista = append(lista, r)
		nBase10 = q
		if q == 0 {
			break
		}
	}
	for i := len(lista) - 1; i >= 0; i-- {
		binario += lista[i]
	}
	fmt.Printf("O número em binário é %v\n", binario)
}
