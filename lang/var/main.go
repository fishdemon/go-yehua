package main

import (
	"fmt"
)

// 常量
// 显示类型定义
const IP string = "localhost"
// 隐式类型定义
const PORT = 7575
// 枚举类定义
const (
	YELLLO = 1
	RED = 2
	GREEN = 3
)
// iota 自增常量
const (
	a = iota   //0
	b          //1s
	c          //2
	d = "ha"   //独立值，iota += 1
	e          //"ha"   iota += 1
	f = 100    //iota +=1
	g          //100  iota +=1
	h = iota   //7,恢复计数
	i          //8
)

// 全局变量
// 不可以使用 := 来声明
var name = "allen"
var colors = []string {"yello","red","blue","black"}


func main() {
	fmt.Println(IP)
	fmt.Println(PORT)
	fmt.Println(YELLLO + RED + GREEN)
	fmt.Println(a, b, c, d, e, f, g, h, i)

	var i int			// 0
	var f float64		// 0
	var b bool			// false
	var s string		// "" 空字符串
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	var p *int
	var arr []int
	var ma map[string]int
	var ch chan int
	var fun func(string) int
	var err error // error 是接口
	fmt.Println(p, arr, ma, ch, fun, err)
	fmt.Println(p == nil)
	fmt.Println(arr == nil)
	fmt.Println(ma == nil)
}

