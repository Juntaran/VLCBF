package General

import "fmt"

// 初始化BloomFilter和索引表
func InitBloom(myBloom []int, indexTable []int, m int) ([]int, []int) {
	for i:=0; i<m; i++ {
		myBloom = append(myBloom, 0)
		indexTable = append(indexTable, i)
	}
	fmt.Printf("length:%v\n", len(myBloom))
	//fmt.Println(myBloom)
	return myBloom, indexTable
}

// 初始化BloomFilter和100索引表
func InitBloom100(myBloom []int, indexTable []int, m int) ([]int, []int) {
	for i:=0; i<m; i++ {
		myBloom = append(myBloom, 0)
	}
	// 100索引表存储的是 indexTable[i]为第(100*i+1)个0所在的位置
	for i:=0; i<m/10; i++ {
		indexTable = append(indexTable, 10 * i)
	}

	//fmt.Printf("length:%v\n", len(myBloom))
	//fmt.Println("IndexTable100:", indexTable)
	return myBloom, indexTable
}