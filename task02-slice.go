package homework

import "fmt"

func reverse(input []int64) (result []int64) {
	//Place your code here

	//copy(input, result)
	result = input
	for i := len(result)/2 - 1; i >= 0; i-- {
		up := len(result) - 1 - i
		result[i], result[up] = result[up], result[i]
	}
	fmt.Println(result)
	return
}
