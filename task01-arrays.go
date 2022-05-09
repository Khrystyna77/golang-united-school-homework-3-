package homework

import "fmt"

func average(input [15]float32) (result float32) {
	//Place your code here
	//Place your code here
	//A := input
	lenmy := len(input)
	var sum float32 = 0

	for i := 0; i < lenmy; i++ {
		sum += input[i]
	}

	result = sum / float32(lenmy)
	fmt.Println(result)
	//no change
	return result
}
