package main

import "fmt"

func potencia(numero, expoente int) int {
	resultadoo := 1
	for i := 0; i < expoente; i++ {
		resultadoo *= numero
	}
	return resultadoo
}

var deno int
var s1, s2, sTot float64

func main() {
	deno = 15
	expo := 0
	for i := 1; i <= 15; i++ {
		if i%2 == 0 {
			s1 += float64((potencia(2, expo)) / (deno * deno))
		} else {
			s2 += float64(potencia(2, expo) / (deno * deno))
		}
		expo += 1
		deno -= 1
	}
	sTot = s2 - s1
	fmt.Println(sTot)
}
