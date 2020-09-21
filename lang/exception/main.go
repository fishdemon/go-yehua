package main

import (
	"fmt"
	"time"
)

func shutdownExec()  {
	panic("j 不可以为 0")
}

// 在其他语言中，通过底层抛出异常，上层通过 try/catch 来捕获并处理异常
// 但是在 go 语言中没有异常系统，它使用 panic 触发宕机（类似于其他语言的抛出异常）
// 然后可以通过 defer 和 recover 配合实现错误捕捉与恢复
// recover 的宕机恢复机制就相当于 try/catch
func recoverExec()  {
	// 延迟处理
	defer func() {
		// 获取 panic 传递的错误信息
		if err := recover(); err != nil {
			fmt.Println("recover: ", err)
		}
		fmt.Println("执行错误处理逻辑...")

		// 在这里面也可以继续向外抛出 panic
	}()

	service(-1)
}

func service(num int) {
	defer func() {
		// 获取 panic 传递的错误信息
		if err := recover(); err != nil {
			fmt.Println("recover: ", err)
		}
		fmt.Println("通知其他服务")
	}()

	if num <= 0 {
		// 主动抛出错误，宕机
		panic("num 必须为正整数")
	}

	// 发生 panic 后，这里不会被执行
	fmt.Println("没有发生panic")
}

func main() {
	fmt.Println("start")
	defer func() {
		// 获取 panic 传递的错误信息
		if err := recover(); err != nil {
			fmt.Println("recover: ", err)
		}
		fmt.Println("通知其他服务")
	}()
	//service(-1)
	//go shutdownExec()
	//go recoverExec()
	go service(-1)
	//fmt.Println("end")

	for true {
		time.Sleep(time.Second *2)
	}
}
