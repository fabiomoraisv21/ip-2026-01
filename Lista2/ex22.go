// Exercicío 22

package main

import "fmt"

func main() {
	var matricula int
	var horasExtras float64

	fmt.Print("Digite a matrícula do funcionário: ")
	fmt.Scanln(&matricula)
	fmt.Print("Digite a quantidade de horas-extras trabalhadas: ")
	fmt.Scanln(&horasExtras)

	const salarioMinimo = 788.00
	const valorHoraExtra = 10.00

	salarioHoraExtra := horasExtras * valorHoraExtra
	salarioBruto := 3*salarioMinimo + salarioHoraExtra

	var descontoINSS float64
	var descontoIR float64

	if salarioBruto > 1500.00 {
		descontoINSS = 0.12 * salarioBruto
	}
	if salarioBruto > 2000.00 {
		descontoIR = 0.20 * salarioBruto
	}
	salarioLiquido := salarioBruto - descontoINSS - descontoIR

	fmt.Printf("Matrícula do funcionário: %d\n", matricula)
	fmt.Printf("Salário bruto: %.2f\n", salarioBruto)
	fmt.Printf("Desconto INSS: %.2f\n", descontoINSS)
	fmt.Printf("Desconto Imposto de Renda: %.2f\n", descontoIR)
	fmt.Printf("Salário líquido: %.2f\n", salarioLiquido)

}
