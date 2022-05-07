package homework

import "fmt"

func average(input [15]float32) (result float32) {
	//Place your code here
	A := input
	lenmy := len(A)
	var sum float32 = 0

	for i := 0; i < lenmy; i++ {
		sum += A[i]
	}
	//fmt.Println(sum / float32(lenmy))
	result = sum / float32(lenmy)
	fmt.Println(result)
	//no change
	return result
}
