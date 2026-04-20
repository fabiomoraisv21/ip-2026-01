package main

import "fmt"

var termo1, termo2, qntdTermos, novoTermo int

func main() {
	fetucine := []int{}
	fmt.Println("Digite os dois primeiros termos da sequência:")
	fmt.Scan(&termo1, &termo2)
	fmt.Println("Digite quantos termos quer na sequência:")
	fmt.Scan(&qntdTermos)
	fetucine = append(fetucine, termo1, termo2)
	for i := 3; i <= qntdTermos; i++ {
		if i%2 == 0 {
			novoTermo = termo2 - termo1
		} else {
			novoTermo = termo1 + termo2
		}
		fetucine = append(fetucine, novoTermo)
		termo2, termo1 = novoTermo, termo2
	}
	fmt.Println(fetucine)

}
