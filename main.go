/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/3/29 06:35
 */

package main

import (
	"fmt"
	"time"
)

var m = 300000 // BloomFilter里最大存储量

func initBloom(myBloom []int, m int) []int {
	// 初始化BloomFilter
	for i := 0; i < m; i++ {
		myBloom = append(myBloom, 0)
	}
	return myBloom
}

func insertEle(myBloom []int, index int) []int {
	// 插入元素
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
	return myBloom
}

func queryEle(myBloom []int, index int) int {
	// 查找元素
	var count = 0
	for i := 0; i < len(myBloom); i++ {
		if myBloom[i] == 0 {
			count++
		}
		if count == index {
			if myBloom[i+1] == 1 {
				//fmt.Println("get!")
				return 1
			}
		}
	}
	return 0
}

func deleteEle(myBloom []int, index int) []int {
	// 删除元素
	var count = 0
	for i := 0; i < len(myBloom); i++ {
		if myBloom[i] == 0 {
			count++
		}
		if count == index {
			if myBloom[i+1] == 1 {
				myBloom = append(myBloom[:i+1], myBloom[i+2:]...)
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

func SDBMHash(word string) int {
	var i = 0
	var hash int
	for {
		if i >= len(word) {
			break
		}
		hash = int(word[i]) + (hash << 6) + (hash << 16) - hash
		i++
	}
	return (hash & 0x7FFFFFFF / 3000)
}

func RSHash(word string) int {
	var i = 0
	var hash int
	var a = 63689
	var b = 378661
	for {
		if i >= len(word) {
			break
		}
		hash = hash*a + int(word[i])
		i++
		a *= b
	}
	return (hash & 0x7FFFFFFF / 3000)
}

func JSHash(word string) int {
	var i = 0
	var hash = 1315423911
	for {
		if i >= len(word) {
			break
		}
		hash ^= ((hash << 5) + int(word[i]) + (hash >> 2))
		i++
	}
	return (hash & 0x7FFFFFFF / 3000)
}

func insertBloom(myBloom []int, word string) []int {
	// 插入一条规则
	var h1 = SDBMHash(word)
	var h2 = JSHash(word)
	myBloom = insertEle(myBloom, h1)
	myBloom = insertEle(myBloom, h2)
	return myBloom
}

func deleteBloom(myBloom []int, word string) []int {
	// 删除一条规则
	var h1 = SDBMHash(word)
	var h2 = JSHash(word)
	myBloom = deleteEle(myBloom, h1)
	myBloom = deleteEle(myBloom, h2)
	return myBloom
}

func queryBloom(myBloom []int, word string) int {
	// 查询一条规则
	var h1 = SDBMHash(word)
	var h2 = JSHash(word)
	var result = queryEle(myBloom, h1) & queryEle(myBloom, h2)
	return result
}

func main() {
	var myBloom []int
	fmt.Printf("length:%v\n", len(myBloom))

	myBloom = initBloom(myBloom, m)
	fmt.Printf("length:%v\n", len(myBloom))

	/* 一个小测试 */
	//myBloom = insertEle(myBloom, 3)
	//fmt.Printf("length:%v\n", len(myBloom))
	//fmt.Println(myBloom)
	//fmt.Println(queryEle(myBloom, 3))
	//myBloom = deleteEle(myBloom, 3)
	//fmt.Printf("length:%v\n", len(myBloom))
	//fmt.Println(myBloom)
	//fmt.Println(queryEle(myBloom, 3))
	/* 小测试结束 */

	var word = "hello"

	var start_time time.Time
	var during_time time.Duration

	start_time = time.Now()
	myBloom = insertBloom(myBloom, word)
	during_time = time.Since(start_time)
	fmt.Println("during time:", during_time)

	start_time = time.Now()
	myBloom = deleteBloom(myBloom, word)
	during_time = time.Since(start_time)
	fmt.Println("during time:", during_time)

	start_time = time.Now()
	queryNTime(myBloom, word, 1)
	during_time = time.Since(start_time)
	fmt.Println("during time:", during_time)

}
