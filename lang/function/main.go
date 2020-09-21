/**
 * 函数的使用, type func() , 有点像闭包
 * 函数与方法是完全不一样的概念, 在go语言中有特别多的使用，而且不好理解
 *
 * 1. 使用 type 定义一个函数
 * 2. 使用 type 定义一个接口，然后函数实现这个接口
 * 3. 定义一个中间方法，参数为这个接口
 * 4. 定义一个主方法，目的是将任意方法强制转化为这个函数，因此参数为方法，调用中间方法，并在调用前将方法强转为函数
 * 5. 在外部就可以调用主方法，并传入用户自定义的回调方法来实现用户的控制
 *
 * http 的 HandleFunc()
 * grpc 的 middleware 注册
 */
package main

import "fmt"

// 回调方法, 可以以参数传递到函数中
// 这个方法名可以随意定义，但是参数必须与函数一直
func callback(name string)  {
	fmt.Println("callback: ", name)
}

// 注册函数,主方法，供外部调用的方法
// 把普通方法强制转换成type定义的函数
func PrintName(name string, cb func(string))  {
	namePrinterHandler(name, NamePrinter(cb))
}

// 中间方法, 方法转换
func namePrinterHandler(name string, h Handler)  {
	h.handle(name)
}

// 定义函数
type NamePrinter func(name string)

// 定义接口
type Handler interface {
	handle(string)
}

// 函数实现 Handler 接口, 实现是由函数来处理
// 从这里可以看出函数与方法是完全不一样的，函数是一种 type，是可以实现接口的
func (f NamePrinter) handle(name string)  {
	f(name)   //使用函数来处理
}

func main() {
	PrintName("allen", callback)
}
