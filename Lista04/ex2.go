package main

import "fmt"

func main() {
	arr1 := make([]int, 10)
	arr2 := make([]int, 5)

	fmt.Println("Primeiro vetor")
	for i := range arr1 {
		fmt.Scan(&arr1[i])
	}

	fmt.Println("Segundo vetor")
	for j := range arr2 {
		fmt.Scan(&arr2[j])
	}

	soma1 := 0
	for v := range arr2 {
		soma1 += arr2[v]
	}

	var arr11 []int
	var arr22 []int

	for j := 0; j < 10; j++ {
		if arr1[j]%2 == 0 {
			arr11 = append(arr11, soma1+arr1[j])
		} else {
			arr22 = append(arr22, soma1+arr1[j])
		}
	}

	fmt.Println(arr11)
	fmt.Println(arr22)

}
