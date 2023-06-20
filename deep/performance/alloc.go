package performance

func NoPreAlloc(size int) []int {
	data := make([]int, 0)
	for k := 0; k < size; k ++ {
		data = append(data, k)
	}
	return data
}

func PreAlloc(size int) []int {
	data := make([]int, 0, size)
	for k := 0; k < size; k ++ {
		data = append(data, k)
	}
	return data
}