package main

import "fmt"

var termos, num, den int
var soma1, soma2, res float64

func main() {
	fmt.Println("Quer quantos termos da sequência?")
	fmt.Scan(&termos)
	num = 1000
	den = 1
	for i := 1; i <= termos; i++ {
		if num%2 == 0 {
			soma1 += float64(num / den)
		} else {
			soma2 += float64(num / den)
		}
		res = soma1 - soma2
		num -= 3
		den += 1
	}
	fmt.Println("O resultado é", res)
}
