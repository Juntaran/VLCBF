/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/3/29 07:14
 */

package main

import (
	"VLCBF_ALL/General"
	"VLCBF_ALL/Hash"
	"fmt"
	"time"
	//"VLCBF_ALL/Test"
)

var M = 300000               // BloomFilter里最大存储量
var indexTable []int         // 索引表存储了每个0的位置，index[i]为第i+1个0所在的位置
var indexTable100 []int      // 100索引表，index[i]为第(100*i+1)个0所在的位置
var check_array [2400000]int // 校验数组，共300000位，每位占4个

/*******************************************************************/
// VLCBF操作
func insertVLCBF(myBloom []int, word string) []int {
	// 插入一条规则
	var h1 = Hash.SDBMHash(word)
	var h2 = Hash.JSHash(word)
	myBloom, indexTable = General.InsertEleIndex100(myBloom, indexTable, h1, M)
	myBloom, indexTable = General.InsertEleIndex100(myBloom, indexTable, h2, M)

	return myBloom
}

func deleteVLCBF(myBloom []int, word string) []int {
	// 删除一条规则
	if queryVLCBF(myBloom, word) == 0 {
		return myBloom
	}
	var h1 = Hash.SDBMHash(word)
	var h2 = Hash.JSHash(word)
	myBloom, indexTable = General.DeleteEleIndex100(myBloom, indexTable, h1, M)
	myBloom, indexTable = General.DeleteEleIndex100(myBloom, indexTable, h2, M)

	return myBloom
}

func queryVLCBF(myBloom []int, word string) int {
	// 查询一条规则
	var h1 = Hash.SDBMHash(word)
	var h2 = Hash.JSHash(word)

	var res1 = General.QueryEleIndex100(myBloom, indexTable, h1)
	var res2 = General.QueryEleIndex100(myBloom, indexTable, h2)

	if res1 == 0 || res2 == 0 {
		// 不在集合中
		return 0
	}
	return 1
}

func updateVLCBF(myBloom []int, word string, count int) {
	// 更新删除操作
	for i := 0; i < count/4; i++ {
		myBloom = insertVLCBF(myBloom, word)
		myBloom = deleteVLCBF(myBloom, word)
	}
	for i := 0; i < count/4; i++ {
		myBloom = deleteVLCBF(myBloom, word)
	}
	for i := 0; i < count/4; i++ {
		myBloom = insertVLCBF(myBloom, word)
	}
}

func queryNTimeVLCBF(myBloom []int, word string, count int) {
	// 查询N次
	for i := 0; i < count; i++ {
		_ = queryVLCBF(myBloom, word)
	}
}

/*******************************************************************/

/*******************************************************************/
// VLCBF_V操作
func insertVLCBF_V(myBloom []int, word string) []int {
	// 插入一条规则
	var h1 = Hash.SDBMHash(word)
	var h2 = Hash.JSHash(word)

	myBloom, indexTable = General.InsertEleIndex100(myBloom, indexTable, h1, M)
	myBloom, indexTable = General.InsertEleIndex100(myBloom, indexTable, h2, M)

	// 校验数组修改
	var checkpos = Hash.RSHash(word)
	//fmt.Println("checkpos:", checkpos)
	A, B, C, D := 0, 0, 0, 0
	if checkpos%2 != 0 {
		A = 1
	}
	if (checkpos/2)%2 != 0 {
		B = 1
	}
	if (checkpos/4)%2 != 0 {
		C = 1
	}
	if (checkpos/8)%2 != 0 {
		D = 1
	}
	check_array[checkpos*4] = check_array[checkpos*4] ^ A
	check_array[checkpos*4+1] = check_array[checkpos*4+1] ^ B
	check_array[checkpos*4+2] = check_array[checkpos*4+2] ^ C
	check_array[checkpos*4+3] = check_array[checkpos*4+3] ^ D

	return myBloom
}

func deleteVLCBF_V(myBloom []int, word string) []int {
	// 删除一条规则
	if queryVLCBF_V(myBloom, word) == 0 {
		return myBloom
	}
	var h1 = Hash.SDBMHash(word)
	var h2 = Hash.JSHash(word)
	myBloom, indexTable = General.DeleteEleIndex100(myBloom, indexTable, h1, M)
	myBloom, indexTable = General.DeleteEleIndex100(myBloom, indexTable, h2, M)

	// 校验数组修改
	var checkpos = Hash.RSHash(word)
	A, B, C, D := 0, 0, 0, 0
	if checkpos%2 != 0 {
		A = 1
	}
	if (checkpos/2)%2 != 0 {
		B = 1
	}
	if (checkpos/4)%2 != 0 {
		C = 1
	}
	if (checkpos/8)%2 != 0 {
		D = 1
	}
	check_array[checkpos*4] = check_array[checkpos*4] ^ A
	check_array[checkpos*4+1] = check_array[checkpos*4+1] ^ B
	check_array[checkpos*4+2] = check_array[checkpos*4+2] ^ C
	check_array[checkpos*4+3] = check_array[checkpos*4+3] ^ D

	return myBloom
}

func queryVLCBF_V(myBloom []int, word string) int {
	// 查询一条规则
	var h1 = Hash.SDBMHash(word)
	var h2 = Hash.JSHash(word)
	var res1 = General.QueryEleIndex100(myBloom, indexTable, h1)
	var res2 = General.QueryEleIndex100(myBloom, indexTable, h2)

	if res1 == 0 || res2 == 0 {
		// 不在集合中
		return 0
	}

	// 校验
	if res1 == 1 && res2 == 1 { // 说明前两个哈希函数找到了
		var checkpos = Hash.RSHash(word) // 校验哈希函数算出来连续4位的值ABCD
		A, B, C, D := 0, 0, 0, 0
		if checkpos%2 != 0 {
			A = 1
		}
		if (checkpos/2)%2 != 0 {
			B = 1
		}
		if (checkpos/4)%2 != 0 {
			C = 1
		}
		if (checkpos/8)%2 != 0 {
			D = 1
		}
		if check_array[checkpos*4] == A && check_array[checkpos*4+1] == B && check_array[checkpos*4+2] == C && check_array[checkpos*4+3] == D {
			// 校验位确定它在集合中
			return 1
		}
	}
	//fmt.Println("不确定")
	return 2
}

func updateVLCBF_V(myBloom []int, word string, count int) {
	// 更新删除操作
	for i := 0; i < count/4; i++ {
		myBloom = insertVLCBF_V(myBloom, word)
		myBloom = deleteVLCBF_V(myBloom, word)
	}
	for i := 0; i < count/4; i++ {
		myBloom = deleteVLCBF_V(myBloom, word)
	}
	for i := 0; i < count/4; i++ {
		myBloom = insertVLCBF_V(myBloom, word)
	}
}

func queryNTimeVLCBF_V(myBloom []int, word string, count int) {
	// 查询N次
	for i := 0; i < count; i++ {
		_ = queryVLCBF_V(myBloom, word)
	}
}

/*******************************************************************/

func main() {
	var myBloom []int // 布隆过滤器
	fmt.Printf("length:%v\n", len(myBloom))

	myBloom, indexTable100 = General.InitBloom100(myBloom, indexTable100, M)
	fmt.Println(indexTable100)
	fmt.Println("length:", len(myBloom))

	/* 一个小测试 */
	//Test.IndexTest(myBloom, indexTable, M)
	//Test.NoneTest(myBloom, M)
	//Test.Index100Test(myBloom, indexTable100, M)
	/* 小测试结束 */

	/*******************************************************************/
	// 100索引
	var word = "hello"
	var TIMES = 100000
	var start_time1, start_time2, start_time3, start_time4 time.Time
	var during_time1, during_time2, during_time3, during_time4 time.Duration

	/*******************************************************************/
	// VLCBF
	fmt.Println("*******************************************")
	fmt.Println("VLCBF:")
	fmt.Println("Times:", TIMES)

	start_time1 = time.Now()
	updateVLCBF(myBloom, word, TIMES/10)
	during_time1 = time.Since(start_time1)
	fmt.Println("Update during time:", during_time1)

	start_time2 = time.Now()
	queryNTimeVLCBF(myBloom, word, TIMES)
	during_time2 = time.Since(start_time2)
	fmt.Println("Query during time:", during_time2)
	fmt.Println("*******************************************")

	/*******************************************************************/
	// VLCBF_V
	fmt.Println("*******************************************")
	fmt.Println("VLCBF_V:")
	fmt.Println("Times:", TIMES)

	General.ResetBloom100(myBloom, indexTable, check_array, M)
	start_time3 = time.Now()
	updateVLCBF_V(myBloom, word, TIMES/10)
	during_time3 = time.Since(start_time3)
	fmt.Println("Update during time:", during_time3)

	start_time4 = time.Now()
	queryNTimeVLCBF_V(myBloom, word, TIMES)
	during_time4 = time.Since(start_time4)
	fmt.Println("Query during time:", during_time4)
	fmt.Println("*******************************************")
}
