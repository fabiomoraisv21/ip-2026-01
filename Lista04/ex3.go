
package main

import "fmt"

func main() {

	arr1 := make([]int, 10)

	fmt.Println("Digite 10 números inteiros:")
	for v := range arr1 {
		fmt.Scan(&arr1[v])
	}

	soma1 := 0
	var arrP []int
	var arrI []int

	for i := 0; i < 10; i++ {
		if arr1[i]%2 == 0 {
			arrP = append(arrP, arr1[i])
			soma1 = soma1 + arr1[i]
		} else {
			arrI = append(arrI, arr1[i])
		}
	}
	fmt.Println(arrP)
	fmt.Println(soma1)
	fmt.Println(arrI)
	fmt.Print(len(arrI))
}
