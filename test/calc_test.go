package main

import (
	"fmt"
	"os"
	"testing"
)

// go test 会全部运行
// go test -v 会全部运行, -v 参数会显示每个用例的测试结果，另外 -cover 参数可以查看覆盖率。

// 前置函数
func setup()  {
	fmt.Println("test start...")
}
// 后置函数
func teardown()  {
	fmt.Println("test finish...")
}
// 配置前置与后置函数, 这样会对当前文件中所有的测试用例生效
func TestMain(m *testing.M)  {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

// go test -run TestAdd -v
func TestAdd(t *testing.T) {
	fmt.Println(Add(1,2))
}

// go test -run TestSub -v
func TestSub(t *testing.T) {
	fmt.Println(Sub(2,1))
}

// 子测试，可以根据不同场景进行测试
func TestMulti(t *testing.T)  {
	
	// go test -run TestMulti/add -v
	t.Run("pos", func(t *testing.T) {
		fmt.Println(Add(1,2))
	})

	// go test -run TestMulti/sub -v
	t.Run("neg", func(t *testing.T) {
		fmt.Println(Add(2,1))
	})
}

// 对于多场景的测试，更推荐下面用法
func TestMulti2(t *testing.T)  {
	cases := []struct{
		Name 	string
		A 		int32
		B 		int32
	}{
		{"pos", 1, 2},
		{"neg", -2, -1},
	}

	for _, item := range cases {
		t.Run(item.Name, func(t *testing.T) {
			fmt.Println(Add(item.A, item.B))
		})
	}
}

