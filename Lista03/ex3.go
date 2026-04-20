package main

import "fmt"

var salarioCarlos, salarioJoao float64
var meses int

func main() {
	fmt.Println("Salário de Carlos:")
	fmt.Scan(&salarioCarlos)
	salarioJoao = salarioCarlos / 3
	for i := 1; i < 1000; i++ {
		salarioCarlos *= 1.02
		salarioJoao *= 1.05
		if salarioJoao >= salarioCarlos {
			meses = i
			break
		}
	}
	fmt.Printf("O salário de João ultrapassou o de Carlos em %v meses\n", meses)
}
