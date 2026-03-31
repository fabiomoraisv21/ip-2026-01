// 3. Composição
package main

import "fmt"

func main() {
	var n1, n2, n3 int

	fmt.Println("Escreva 3 números:")
	fmt.Scan(&n1, &n2, &n3)

	if n1 > 9 || n2 > 9 || n3 > 9 {
		fmt.Println("DIGITO INVALIDO")
	} else {
		numero := n1*100 + n2*10 + n3
		fmt.Println(numero, ",", numero*numero)
	}
}
