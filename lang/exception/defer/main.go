/**
 * defer 的妙用
 * defer 与 return 的执行顺序你真的知道么
 */
package main

import (
	"fmt"
	"os"
)

// 不是 10 , 是 11
func defer1() (res int) {
	defer func() {
		// 2. res = 10 + 1 = 11
		res ++
	}()
	// 1. res = 10
	return 10
}

// 不是15, 是10
func defer2() (res int) {
	// 1. sb = 10
	sb := 10
	defer func() {
		// 3. sb = 15, 但是 res = 10
		sb += 5
	}()
	// 2. res = sb = 10
	return sb
}

// 不是12，是10
func defer3() (res int) {
	// 1. res = 2
	res = 2
	defer func(res int) {
		// 内部res为形参，不影响外边的值 res=2+2=4
		res += 2
		fmt.Println("内部 res ", res)   // 4
	}(res)		// defer 参数的值是在声明的时候确定的，也就是只有 defer 之前的语句会影响这个值, res = 2
	// 2. res = 10
	return 10
}

// 1. 同一方法中的多个defer, 是从后往前执行的
// 2. defer 声明时有两种对外部变量的引用方式
//   (1) 参数传递：在defer声明时，即将值传递给defer,并缓存起来，调用defer的时候使用缓存值进行计算
//   (2) 直接引用：根据上下文确定当前值，作用域与外部变量相同
func deferTest() {
	fish := 0
	// 直接引用(闭包)
	defer func() {
		fmt.Println("d1: ", fish)	// 3
	}()
	// 参数传递
	defer fmt.Println("d2: ", fish) // 0
	// 参数传递(闭包)
	defer func(fish int) {
		fish += 2						// fish 只作用于内部
		fmt.Println("d3: ", fish)	// 2
	}(fish)							    // 声明时传递, fish = 0
	// 直接引用(闭包)
	defer func() {
		fmt.Println("d4: ", fish)	// 1
		fish +=2						// fish = 3, 作用域与外部的相同
	}()
	fish++
}

func OpenFile()  {
	file, err := os.Open("/home/demo/info.txt")
	if err != nil {
		fmt.Println(err)
	}

	// 这里将 defer 放在 err 判断的后面，而不是 os.Open() 的后面
	// 因为若 err != nil ，文件打开是失败的, 没必要释放
	// 若 err != nil, file 有可能为 nil ,这时候释放资源可能会导致程序崩溃
	defer file.Close()
}


func main() {
	fmt.Println(defer1())
	fmt.Println(defer2())
	fmt.Println(defer3())

	deferTest()
}
