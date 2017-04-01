package Test

import (
	"fmt"
	"VLCBF_ALL/General"
)

// 100索引测试
func Index100Test(myBloom []int, indexTable []int, m int)  {
	fmt.Println("/************************************/")
	fmt.Println("100索引测试")
	//fmt.Println(myBloom)
	myBloom, indexTable = General.InsertEleIndex100(myBloom, indexTable, 103, m)
	fmt.Printf("length:%v\n", len(myBloom))
	//fmt.Println(myBloom)
	fmt.Println(General.QueryEleIndex100(myBloom, indexTable, 103))
	fmt.Println(General.QueryEleIndex100(myBloom, indexTable, 7))
	myBloom, indexTable = General.DeleteEleIndex100(myBloom, indexTable, 7, m)
	//fmt.Println(myBloom)
	myBloom, indexTable = General.DeleteEleIndex100(myBloom, indexTable, 103, m)
	fmt.Printf("length:%v\n", len(myBloom))
	fmt.Println(General.QueryEleIndex100(myBloom, indexTable, 3))
	//fmt.Println(myBloom)
	fmt.Println("/************************************/")
}

// 有索引测试
func IndexTest(myBloom []int, indexTable []int, m int)  {
	fmt.Println("/************************************/")
	fmt.Println("有索引测试")
	fmt.Println(myBloom)
	myBloom, indexTable = General.InsertEleIndex(myBloom, indexTable, 3, m)
	fmt.Printf("length:%v\n", len(myBloom))
	fmt.Println(myBloom)
	fmt.Println(General.QueryEleIndex(myBloom, indexTable, 3))
	fmt.Println(General.QueryEleIndex(myBloom, indexTable, 7))
	myBloom, indexTable = General.DeleteEleIndex(myBloom, indexTable, 7, m)
	fmt.Println(myBloom)
	myBloom, indexTable = General.DeleteEleIndex(myBloom, indexTable, 3, m)
	fmt.Printf("length:%v\n", len(myBloom))
	fmt.Println(General.QueryEleIndex(myBloom, indexTable, 3))
	fmt.Println(myBloom)
	fmt.Println("/************************************/")
}

// 无索引测试
func NoneTest(myBloom []int, m int)  {
	fmt.Println("/************************************/")
	fmt.Println("无索引测试")
	fmt.Println(myBloom)
	myBloom = General.InsertEle(myBloom, 7)
	fmt.Printf("length:%v\n", len(myBloom))
	fmt.Println(myBloom)
	fmt.Println(General.QueryEle(myBloom, 7))
	fmt.Println(General.QueryEle(myBloom, 4))
	myBloom = General.DeleteEle(myBloom, 3, m)
	fmt.Println(myBloom)
	myBloom = General.DeleteEle(myBloom, 7, m)
	fmt.Printf("length:%v\n", len(myBloom))
	fmt.Println(General.QueryEle(myBloom, 7))
	fmt.Println(myBloom)
	fmt.Println("/************************************/")
}
