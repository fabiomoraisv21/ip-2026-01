package main

import "fmt"

var alunos, aprovados, exame, reprovados int
var nota1, nota2, mediaAluno, somaSala, mediaSala float64

func main() {
	fmt.Println("Quantos alunos?")
	fmt.Scan(&alunos)
	for i := 0; i < alunos; i++ {
		fmt.Println("Digite a nota 1:")
		fmt.Scan(&nota1)
		fmt.Println("Digite a nota 2:")
		fmt.Scan(&nota2)
		mediaAluno = (nota1 + nota2) / 2
		fmt.Println("A média do aluno foi de", mediaAluno)
		if mediaAluno <= 3 {
			fmt.Println("Aluno reprovado")
			reprovados += 1
		} else if mediaAluno > 3 && mediaAluno < 7 {
			fmt.Println("Exame")
			exame += 1
		} else {
			fmt.Println("Aluno aprovado")
			aprovados += 1
		}
		somaSala += mediaAluno
	}
	mediaSala = somaSala / float64(alunos)
	fmt.Printf("%v alunos aprovados\n", aprovados)
	fmt.Printf("%v alunos em exame\n", exame)
	fmt.Printf("%v alunos reprovados\n", reprovados)
	fmt.Printf("A média da sala foi de %v\n", mediaSala)
}
