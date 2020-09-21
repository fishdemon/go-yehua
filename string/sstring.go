package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// 可以定义多行字符串（这里面的换行符是保留的，打印出来也是3行）
var multiStr1 = `this is 1 line
this is 2 line
this is 3 line`

// 这里面的tab缩进将被保留
var multiStr2 = `this is 1 line
	this is 2 line
	this is 3 line`


func printString()  {
	str := "hello world"

	// 字符串中的字符打印
	for i:=0 ; i<len(str); i++ {
		fmt.Printf("%c", str[i])
	}
}

func int2String()  {
	str := "hello world"
	fmt.Println("str的长度：" + strconv.Itoa(len(str)))

	var i1 int32 = 1021
	var i2 int64 = int64(i1)
	var f float32 = float32(i1)
	fmt.Printf("i1 : %d, i2 : %d", i1, i2)
	fmt.Printf("%f", f)
}

func testStrings() {

}

// 字符串连接方式
// 1. + 连接比较低效(适合少量)
// 2. bytes.Buffer 来连接（适合大量连接）
// 3. string.Join 适合连接字符串数组
func connectStr()  {

	// 与 java 中的 StringBuffer 目的一样，都是为了更快的连接字符串
	var b bytes.Buffer
	// 连接字符串
	b.WriteString("hello")
	// 连接字节数组
	b.Write([]byte(" world! "))
	// 连接字节
	b.WriteByte(byte('i'))
	// 连接 unicode 字符
	b.WriteRune(63)
	fmt.Println(b.String())			// hello world! i?


	var strs []string
	for i := 0; i < 10; i++ {
		strs = append(strs, "hello")
	}
	// Join 来连接字符串数组，并可以加分隔符
	res := strings.Join(strs, "?")
	fmt.Println(res)				// hello?hello?hello?hello?hello?hello?hello?hello?hello?hello
}

// 生成随机的字符串
func randString(length int) string {
	// 自定义随机字符, 可以定义成常量
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// 随机数源
	source := rand.NewSource(time.Now().UnixNano())

	b := make([]byte, length)
	for i := 0; i < length; i++ {
		b[i] = charset[source.Int63()%int64(len(charset))]
	}

	return string(b)
}

func main()  {

	//arr := []int{1,2,3,4}
	//fmt.Sprintf("%p", &arr)
	//fmt.Sprintf("%p", unsafe.Pointer(&arr))

	fmt.Printf("test%%s")

	fmt.Println()

	fmt.Printf("%s", time.Now())

	//connectStr()
	fmt.Println(time.Now().Format("2006-01-02"))
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	fmt.Println(yesterday)

	str := ""
	fmt.Println(strings.Split(str, ","))

	colors := []string{"1", "2", "3"}
	fmt.Print(strings.Join(colors, ","))

}

