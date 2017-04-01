package General

import "fmt"

// 100索引表的插入元素操作
func InsertEleIndex100(myBloom []int, indexTable []int, index int, m int) ([]int, []int) {
	// 在第index个0后面插入元素
	//fmt.Println("befor insert:", myBloom, indexTable)
	if index >= m {
		fmt.Println("insert Error")
		fmt.Println("insert index:", index)
		return myBloom, indexTable
	}

	// 更新100索引表
	for i := (index + 10) / 10; i < len(indexTable); i++ {
		indexTable[i]++
	}
	var k = (index+10)/10 - 1
	var count = 0
	for i := indexTable[k]; i < len(myBloom); i++ {
		//fmt.Println("***")
		if myBloom[i] == 0 {
			count++
		}
		if count == (index - k*10) {
			last := append([]int{}, myBloom[i+1:]...)
			myBloom = append(myBloom[0:i+1], 1)
			myBloom = append(myBloom, last...)
			return myBloom, indexTable
			break
		}
	}
	//fmt.Println("after insert:", myBloom, indexTable)
	return myBloom, indexTable
}

// 有索引表的插入元素操作
func InsertEleIndex(myBloom []int, indexTable []int, index int, m int) ([]int, []int) {
	// 在第index个0后面插入元素
	//fmt.Println("befor insert:",myBloom)
	if index >= m {
		fmt.Println("Error")
		return myBloom, indexTable
	}

	// 更新索引表
	for i := index; i < len(indexTable); i++ {
		indexTable[i]++
	}
	last := append([]int{}, myBloom[indexTable[index-1]+1:]...)
	myBloom = append(myBloom[0:indexTable[index-1]+1], 1)
	myBloom = append(myBloom, last...)
	return myBloom, indexTable
}

// 无索引表的插入操作
func InsertEle(myBloom []int, index int) []int {
	//fmt.Println("befor insert:",myBloom)
	if index >= len(myBloom) {
		fmt.Println("Error")
		return myBloom
	}

	var count = 0
	for i := 0; i < len(myBloom); i++ {
		if myBloom[i] == 0 {
			count++
		}
		if count == index {
			last := append([]int{}, myBloom[i+1:]...)
			myBloom = append(myBloom[0:i+1], 1)
			myBloom = append(myBloom, last...)
			return myBloom
			break
		}
	}
	//fmt.Println("after insert:",myBloom)
	return myBloom
}
