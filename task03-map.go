package homework

import (
	"fmt"
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	keys := make([]int, 0, len(input))

	i := 0
	for k := range input {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	//fmt.Println(keys)

	for _, k := range keys {
		fmt.Println(sortMapValues(k))
	}
	return

}
