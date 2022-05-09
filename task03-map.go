package homework

import (
	"fmt"
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	//for i := 0; i < len(input); i++ {
	keys := make([]int, len(input))
	i := 0
	for k := range input {		
		keys[i] = k
		i++
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// To perform the opertion you want
	for _, k := range keys {
		fmt.Println(sortMapValues(k:sortMapValues()))
	}
	return
}
