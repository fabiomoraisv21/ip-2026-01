// 4. Energia
package main

import "fmt"

func main() {
	var salario, kw float64

	fmt.Println("Digite o valor do salário e a quantidade de kw gasto na residência:")
	fmt.Scan(&salario, &kw)

	valorKw := (salario * 0.7) / 100
	total := valorKw * kw
	desconto := total * 0.9

	fmt.Println("Custo por kW: R$", valorKw)
	fmt.Println("Custo do consumo: R$", total)
	fmt.Println("Custo com desconto: R$", desconto)
}
