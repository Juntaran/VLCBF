package General

//import "fmt"

// 100索引表的查询元素操作
func QueryEleIndex100(myBloom []int, indexTable []int, index int) int {

	var k = (index+10)/10 - 1
	var count = 0

	for i := 10 * k; i < len(myBloom); i++ {
		if myBloom[i] == 0 {
			count++
		}
		if count == (index - k) {
			//fmt.Println("Bingo!")
			if myBloom[i+1] == 1 {
				//fmt.Println("get!")
				if myBloom[i+2] != 1 {
					return 1 // 查询到可以使用校验位
				}
				return -1 // 查询到了但是不能使用校验位
			}
		}
	}
	return 0 // 该元素不存在
}

// 有索引表的查询元素操作
func QueryEleIndex(myBloom []int, indexTable []int, index int) int {
	if myBloom[indexTable[index-1]+1] == 1 {
		if myBloom[indexTable[index-1]+2] != 1 {
			return 1 // 查询到可以使用校验位
		}
		return -1 // 查询到了但是不能使用校验位
	}
	return 0 // 该元素不存在
}

// 没有索引表的查询元素操作
func QueryEle(myBloom []int, index int) int {
	var count = 0
	for i := 0; i < len(myBloom); i++ {
		if myBloom[i] == 0 {
			count++
		}
		if count == index {
			//fmt.Println("bingo!", i)
			if myBloom[i+1] == 1 {
				//fmt.Println("get!")
				if myBloom[i+2] != 1 {
					return 1 // 查询到可以使用校验位
				}
				return -1 // 查询到了但是不能使用校验位
			}
		}
	}
	return 0 // 该元素不存在
}
