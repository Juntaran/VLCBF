package General

// 重置BloomFilter、索引表、校验位
func ResetBloom(myBloom []int, indexTable []int, check_array [2400000]int, m int) ([]int, []int, [2400000]int) {
	for i := 0; i < m; i++ {
		myBloom[i] = 0
		indexTable[i] = i
	}
	for i := 0; i < 2400000; i++ {
		check_array[i] = i
	}
	return myBloom, indexTable, check_array
}

// 重置BloomFilter、100索引表、校验位
func ResetBloom100(myBloom []int, indexTable []int, check_array [2400000]int, m int) ([]int, []int, [2400000]int) {
	for i := 0; i < m; i++ {
		myBloom[i] = 0
	}
	for i := 0; i < 2400000; i++ {
		check_array[i] = i
	}
	for i := 0; i < m/10; i++ {
		indexTable = append(indexTable, 10*i)
	}
	return myBloom, indexTable, check_array
}
