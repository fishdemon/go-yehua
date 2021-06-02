package main

import (
	"fmt"
	"strconv"
)

// 保留两位小数
func floatDemo()  {
	f1 := 1.2345
	// 1. 使用格式化成字符串，然后再转换为浮点型
	f, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", f1), 64)
	fmt.Println(f)
	// 2.

}
