package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// 这个字符串含有8个字符（5个英文，一个空格，2个汉子）
	str := "hello 张三"
	// 但这里并不是 8, 而是 12
	fmt.Println(len(str))			// 12

	// 这是怎么回事？与我们的认知完全不一样
	// 其实 len 计算的是底层存储的字节数，一个英文占一个字节，一个中文在utf-8编码下占用3个字节，在 unicode 下占用2个字节
	// go 语言默认使用 utf-8 编码

	// 那么我们如何获取字符串的长度呢？而不是底层所占的字节长度
	// 这时候就可以使用 rune 这个类型了, rune相当于 int32, 但保存的时候 utf-8/unicode 的字符编码值
	fmt.Println(len([]rune(str)))


	// utf-8 包下也提供方法来计算字符长度
	fmt.Println(utf8.RuneCountInString(str))

	// byte 等同于int8，常用来处理ascii字符
	// rune 等同于int32,常用来处理unicode或utf-8字符
}
