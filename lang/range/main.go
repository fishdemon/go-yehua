package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// 数组
// 定义一个固定数组
var names = [2]string{"allen","june"}
// 定义一个固定数组，大小自动读取
var ages = [...]int{23,22,56,23,12}
// 定义一个二维数组，2行3列
var a = [2][3]int {
	{1,2,3},
	{2,3,4},
}

// 切片 动态数组
var colors = []string {"yello","red","blue","black"}

// map 集合
var person = map[string]string{"name":"allen", "age":"28", "sex":"male"}

// 结构体
type User struct {
	id     int32
	name   string
	age    int8
	sex    string
	parent *User			// 嵌套自身
	children []*User		// 结构体数组
}

//
var user = struct {
	id   int32
	name string
	age  int8
	sex  string
}{}

func main() {

	// for 循环遍历数组
	for i := 0; i < len(names); i++ {
		fmt.Print(names[i] + " ")
	}
	fmt.Println()

	// range 迭代遍历
	// k 是index, v是值，数组可以看作是一种特殊类型的hashmap
	for k, v := range ages {
		fmt.Print( strconv.Itoa(k) + ":" + strconv.Itoa(v) + " ")
	}
	fmt.Println()

	// range 迭代二维数组
	for k,v := range a {
		for k1,v1 := range v {
			fmt.Print(k)
			fmt.Print(",")
			fmt.Print(k1)
			fmt.Print(":")
			fmt.Println(v1)
		}
	}
	fmt.Println()

	// for 循环遍历数组
	for i := 0; i < len(colors); i++ {
		fmt.Print(colors[i] + " ")
	}
	fmt.Println()

	// range 迭代遍历切片
	// k 是index, v是值，数组可以看作是一种特殊类型的hashmap
	for k, v := range colors {
		fmt.Print( strconv.Itoa(k) + ":" + v + " ")
	}
	fmt.Println()

	// range 迭代 map
	for k, v:= range person {
		fmt.Print( k + ":" + v + " ")
	}
	fmt.Println()

	u := &User{1, "allen", 28, "male", nil, nil}
	v := reflect.ValueOf(u)
	count := v.NumField()
	for i := 0; i < count; i++ {
		f := v.Field(i)
		switch f.Kind() {
		case reflect.String:
			fmt.Println(f.String())
		case reflect.Int:
			fmt.Println(f.Int())
		}
	}

	// 相当于 while 条件循环
	i := 1
	for i < 3 {
		fmt.Println("")
	}
	// 相当于 while ， 无限循环
	for {
		fmt.Println("")
	}
}

