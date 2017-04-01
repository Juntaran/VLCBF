/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/3/29 07:14
 */

package main

import (
	"VLCBF_ALL/General"
	"VLCBF_ALL/Hash"
	"VLCBF_ALL/ReadFile"
	"fmt"
	"time"
	//"VLCBF_ALL/Test"
)

var M = 300000               // BloomFilter里最大存储量
var indexTable []int         // 索引表存储了每个0的位置，index[i]为第i+1个0所在的位置
var indexTable100 []int      // 100索引表，index[i]为第(100*i+1)个0所在的位置
var check_array [2400000]int // 校验数组，共300000位，每位占4个
var stringSlice []string     // 字符串切片

/*******************************************************************/
// VLCBF操作
func insertVLCBF(myBloom []int, word string) []int {
	// 插入一条规则
	var h1 = Hash.SDBMHash(word)
	var h2 = Hash.JSHash(word)
	myBloom, indexTable100 = General.InsertEleIndex100(myBloom, indexTable100, h1, M)
	myBloom, indexTable100 = General.InsertEleIndex100(myBloom, indexTable100, h2, M)

	return myBloom
}

func deleteVLCBF(myBloom []int, word string) []int {
	// 删除一条规则
	if queryVLCBF(myBloom, word) == 0 {
		return myBloom
	}
	var h1 = Hash.SDBMHash(word)
	var h2 = Hash.JSHash(word)
	myBloom, indexTable100 = General.DeleteEleIndex100(myBloom, indexTable100, h1, M)
	myBloom, indexTable100 = General.DeleteEleIndex100(myBloom, indexTable100, h2, M)
	return myBloom
}

func queryVLCBF(myBloom []int, word string) int {
	// 查询一条规则
	var h1 = Hash.SDBMHash(word)
	var h2 = Hash.JSHash(word)
	var res1 = General.QueryEleIndex100(myBloom, indexTable100, h1)
	var res2 = General.QueryEleIndex100(myBloom, indexTable100, h2)

	if res1 == 0 || res2 == 0 {
		// 不在集合中
		return 0
	}
	return 1
}

func updateVLCBF(myBloom []int, word []string, count int) {
	// 更新删除操作count轮
	for i := 0; i < count; i++ {
		// 每轮更新操作1000次
		for j := 0; j < 250; j++ {
			myBloom = insertVLCBF(myBloom, word[j])
			myBloom = deleteVLCBF(myBloom, word[j])
		}
		for j := 250; j < 500; j++ {
			myBloom = deleteVLCBF(myBloom, word[j])
		}
		for j := 500; j < 750; j++ {
			myBloom = insertVLCBF(myBloom, word[j])
		}
	}
}

func queryNTimeVLCBF(myBloom []int, word []string, count int) {
	// 查询count轮
	for i := 0; i < count; i++ {
		// 每轮查询操作1000次
		for j := 0; j < 1000; j++ {
			_ = queryVLCBF(myBloom, word[j])
		}
	}
}

/*******************************************************************/

/*******************************************************************/
// VLCBF_V操作
func insertVLCBF_V(myBloom []int, word string) []int {
	// 插入一条规则
	var h1 = Hash.SDBMHash(word)
	var h2 = Hash.JSHash(word)
	myBloom, indexTable100 = General.InsertEleIndex100(myBloom, indexTable100, h1, M)
	myBloom, indexTable100 = General.InsertEleIndex100(myBloom, indexTable100, h2, M)

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
	//fmt.Println("insert success")
	return myBloom
}

func deleteVLCBF_V(myBloom []int, word string) []int {
	// 删除一条规则
	if queryVLCBF_V(myBloom, word) == 0 {
		return myBloom
	}
	var h1 = Hash.SDBMHash(word)
	var h2 = Hash.JSHash(word)
	myBloom, indexTable100 = General.DeleteEleIndex100(myBloom, indexTable100, h1, M)
	myBloom, indexTable100 = General.DeleteEleIndex100(myBloom, indexTable100, h2, M)

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
	//fmt.Println("delete success")
	return myBloom
}

func queryVLCBF_V(myBloom []int, word string) int {
	// 查询一条规则
	var h1 = Hash.SDBMHash(word)
	var h2 = Hash.JSHash(word)
	var res1 = General.QueryEleIndex100(myBloom, indexTable100, h1)
	var res2 = General.QueryEleIndex100(myBloom, indexTable100, h2)
	if res1 == 0 || res2 == 0 {
		// 不在集合中
		//fmt.Println("query Fail 1")
		return 0
	}

	// 校验
	if res1 == 1 && res2 == 1 { // 说明前两个哈希函数找到了
		//fmt.Println("start check")
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
			//fmt.Println("check success")
			return 1
		}
		//fmt.Println("check fail")
		return 0
	}
	return 0
}

func updateVLCBF_V(myBloom []int, word []string, count int) {
	// 更新删除操作count轮
	for i := 0; i < count; i++ {
		// 每轮更新操作1000次
		for j := 0; j < 250; j++ {
			myBloom = insertVLCBF_V(myBloom, word[j])
			myBloom = deleteVLCBF_V(myBloom, word[j])
		}
		for j := 250; j < 500; j++ {
			myBloom = deleteVLCBF_V(myBloom, word[j])
		}
		for j := 500; j < 750; j++ {
			myBloom = insertVLCBF_V(myBloom, word[j])
		}
	}
}

func queryNTimeVLCBF_V(myBloom []int, word []string, count int) {
	// 查询count轮
	for i := 0; i < count; i++ {
		// 每轮查询操作1000次
		for j := 0; j < 1000; j++ {
			_ = queryVLCBF_V(myBloom, word[j])
		}
	}
}

/*******************************************************************/

func main() {
	var myBloom []int // 布隆过滤器
	fmt.Printf("length:%v\n", len(myBloom))
	//fmt.Println(check_array[400])
	//myBloom, indexTable = General.InitBloom(myBloom, indexTable, M)
	//fmt.Printf("length:%v\n", len(myBloom))

	myBloom, indexTable100 = General.InitBloom100(myBloom, indexTable100, M)

	fmt.Println(indexTable100)
	fmt.Println("length:", len(myBloom))

	/* 一个小测试 */
	//Test.IndexTest(myBloom, indexTable, M)
	//Test.NoneTest(myBloom, M)
	//Test.Index100Test(myBloom, indexTable100, M)
	/* 小测试结束 */

	stringSlice, _ = ReadFile.ReadFile("ReadFile/test.txt", stringSlice)

	/*******************************************************************/
	// 100索引

	var TIMES = 1
	var start_time time.Time
	var during_time time.Duration

	/*******************************************************************/
	// VLCBF
	fmt.Println("*******************************************")
	fmt.Println("VLCBF:")
	fmt.Println("Times:", TIMES)

	var myBloom_bak = myBloom
	var indexTable100_bak = indexTable100
	var check_array_bak = check_array

	start_time = time.Now()
	updateVLCBF(myBloom, stringSlice, 1)
	during_time = time.Since(start_time)
	fmt.Println("Update during time:", during_time)

	start_time = time.Now()
	queryNTimeVLCBF(myBloom, stringSlice, TIMES)
	during_time = time.Since(start_time)
	fmt.Println("Update during time:", during_time)

	indexTable100 = indexTable100_bak
	check_array = check_array_bak
	myBloom = myBloom_bak
	fmt.Println("*******************************************")
	/*******************************************************************/

	General.ResetBloom100(myBloom, indexTable100, check_array, M)

	/*******************************************************************/
	// VLCBF_V
	fmt.Println("*******************************************")
	fmt.Println("VLCBF_V:")
	fmt.Println("Times:", TIMES)

	var myBloom_bak = myBloom
	var indexTable100_bak = indexTable100
	var check_array_bak = check_array

	start_time = time.Now()
	updateVLCBF_V(myBloom, stringSlice, 1)
	during_time = time.Since(start_time)
	fmt.Println("Update during time:", during_time)

	start_time = time.Now()
	queryNTimeVLCBF_V(myBloom, stringSlice, TIMES)
	during_time = time.Since(start_time)
	fmt.Println("Update during time:", during_time)

	indexTable100 = indexTable100_bak
	check_array = check_array_bak
	myBloom = myBloom_bak
	fmt.Println("*******************************************")
	/*******************************************************************/
}
