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
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Println(sortMapValues[k])
	}
	return
}
