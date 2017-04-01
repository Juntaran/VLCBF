package General

import "fmt"

// 100索引表的删除元素操作
func DeleteEleIndex100(myBloom []int, indexTable []int, index int, m int) ([]int, []int) {
	// 在第index个元素后面删除元素
	//fmt.Println("before delete:", myBloom, indexTable)
	if index > m {
		fmt.Println("Error")
		return myBloom, indexTable
	}

	var k = (index+10)/10 - 1
	var count = 0
	for i := 100 * k; i < len(myBloom); i++ {
		//fmt.Println("***")
		if myBloom[i] == 0 {
			count++
		}
		if count == (index - k) {
			if myBloom[i+1] == 1 {
				myBloom = append(myBloom[:i+1], myBloom[i+2:]...)

				// 更新100索引表
				for i := (index + 10) / 10; i < len(indexTable); i++ {
					indexTable[i]--
				}
				//fmt.Println("after  delete:", myBloom, indexTable)
				return myBloom, indexTable
				break
			} else {
				fmt.Println("Not Found!")
				return myBloom, indexTable
			}
		}
	}

	return myBloom, indexTable
}

// 有索引表的删除元素操作
func DeleteEleIndex(myBloom []int, indexTable []int, index int, m int) ([]int, []int) {
	// 在第index个元素后面删除元素
	//fmt.Println("before delete:", myBloom)
	if index > m {
		fmt.Println("Error")
		return myBloom, indexTable
	}

	// 更新索引表
	for i := index; i < len(indexTable); i++ {
		indexTable[i]--
	}

	if myBloom[indexTable[index-1]+1] == 1 {
		myBloom = append(myBloom[:indexTable[index-1]+1], myBloom[indexTable[index-1]+2:]...)
	} else {
		//fmt.Println("Not Found!")
		return myBloom, indexTable
	}

	//fmt.Println("after delete:", myBloom)
	return myBloom, indexTable
}

// 无索引表的删除元素操作
func DeleteEle(myBloom []int, index int, m int) []int {
	var count = 0
	for i := 0; i < len(myBloom); i++ {
		//fmt.Println("***")
		if myBloom[i] == 0 {
			count++
		}
		if count == index {
			if myBloom[i+1] == 1 {
				myBloom = append(myBloom[:i+1], myBloom[i+2:]...)
				//fmt.Println("after delete:",myBloom)
				//fmt.Println("delete!")
				return myBloom
				break
			} else {
				//fmt.Println("Not Found!")
				return myBloom
			}
		}
	}
	return myBloom
}
