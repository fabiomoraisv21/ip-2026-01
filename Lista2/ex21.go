// Exercicío 21

package main

import "fmt"

func main() {

	var id int
	var nota1, nota2, nota3, mediaExercicios float64

	fmt.Print("Digite o número de identificação do aluno: ")
	fmt.Scanln(&id)
	fmt.Print("Digite a nota da primeira verificação: ")
	fmt.Scanln(&nota1)
	fmt.Print("Digite a nota da segunda verificação: ")
	fmt.Scanln(&nota2)
	fmt.Print("Digite a nota da terceira verificação: ")
	fmt.Scanln(&nota3)
	fmt.Print("Digite a média dos exercícios: ")
	fmt.Scanln(&mediaExercicios)

	mediaAproveitamento := (nota1 + nota2*2 + nota3*3 + mediaExercicios) / 7

	var conceito string
	var mensagem string

	if mediaAproveitamento >= 9.0 && mediaAproveitamento <= 10.0 {
		conceito = "A"
	} else if mediaAproveitamento >= 7.5 && mediaAproveitamento < 9.0 {
		conceito = "B"
	} else if mediaAproveitamento >= 6.0 && mediaAproveitamento < 7.5 {
		conceito = "C"
	} else if mediaAproveitamento >= 4.0 && mediaAproveitamento < 6.0 {
		conceito = "D"
	} else {
		conceito = "E"
	}

	if conceito == "A" || conceito == "B" || conceito == "C" {
		mensagem = "APROVADO"
	} else {
		mensagem = "REPROVADO"
	}

	fmt.Printf("Número do aluno: %d\n", id)
	fmt.Printf("Notas: %.2f, %.2f, %.2f\n", nota1, nota2, nota3)
	fmt.Printf("Média dos exercícios: %.2f\n", mediaExercicios)
	fmt.Printf("Média de aproveitamento: %.2f\n", mediaAproveitamento)
	fmt.Printf("Conceito: %s\n", conceito)
	fmt.Printf("Mensagem: %s\n", mensagem)

}
