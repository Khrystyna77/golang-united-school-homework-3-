package homework

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	keys := make([]int, 0, len(input))

	for k := range input {
	   keys = append(keys, k)
	}
	sort.Ints(keys)
	//fmt.Println(keys)

	for _, k := range keys {
	   fmt.Println(k, sortMapValues[k])
	return

	return
}
