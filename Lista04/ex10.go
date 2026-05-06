package main

import "fmt"

func main() {

	var arr [50]int
	arr[0] = 1
	arr[1] = 1

	for i := 2; i < 50; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	for j := 0; j < 50; j++ {
		fmt.Print(arr[j],"-")
	}
}