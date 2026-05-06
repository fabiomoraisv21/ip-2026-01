

package main

import "fmt"

func main () {

	arr := make([]int, 100)

	arr[0] = 1

	for i := 1; i < 100; i++{
	
		arr[i] = 2 * i + arr[0]
	}

	fmt.Print(arr)
}
