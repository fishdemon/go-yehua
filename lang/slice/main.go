package main

import "fmt"

func main() {
	arr := [3]int{1,2,3}

	// 创建切片
	var s1 []int     			// 创建一个nil切片
	s2 := make([]int, 0, 0)		// 创建一个空切片
	s3 := make([]int, 0, 100)	// 创建一个有容量切片
	s4 := make([]int, 100)		// 创建一个零值切片
	s5 := arr[:]				// 引用数组创建切片

	// 内置函数来处理切片
	fmt.Println(len(s2))		// 切片长度
	fmt.Println(cap(s2))		// 切片容量
	s3 = append(s3, 1)	// 追加内容
	copy(s5, s4)				// 将s4 的内容拷贝到 s5 中

	s1 = append(s1, 2)	// nil 切片可以直接追加
	fmt.Println(s1)
	s2 = append(s2, 2)	// 空切片也可以直接追加
	fmt.Println(s2)

	// 切片的冒号操作
	s := []int{1,2,3,4,5,6,7,8,9,0}
	fmt.Println(s[:])			// 引用一个切片
	fmt.Println(s[:0])			// 清空切片
	fmt.Println(s[:8])			// 截取切片，index 从0到8（不包含8）
	fmt.Println(s[2:3])			// 截取切片
	fmt.Println(s[1:])			// 截取切片
	// [i,l,c] 操作, i<=l<=c
	fmt.Println(s[1:8:9])		// [2 3 4 5 6 7 8]
	fmt.Println(cap(s[1:8:9]))	// 8
	fmt.Println(s[1:10:20])     // 会报错,c 不可以超过当前数组的长度

}
