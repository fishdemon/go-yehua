package main

import (
	"fmt"
	"os"
)

type Person struct {
	Name string
	Age uint8
	Sex string
}

func testPrintf()  {
	var ch byte = 'b'
	fmt.Println(ch)
	ch = 66
	fmt.Println(ch)
	ch = 0x41
	fmt.Println(ch)

	var uch byte = '\u0045'
	var uch1 int = '\u03b9'

	fmt.Printf("%d, %d", uch, uch1) //int
	fmt.Println()
	fmt.Printf("%c, %c", uch, uch1) //字符
	fmt.Println()
	fmt.Printf("%X, %X", uch, uch1) //UTF-8字节
	fmt.Println()
	fmt.Printf("%U, %U", uch, uch1) //UTF-8 code point
	fmt.Println()

	p := Person{"allen", 1, "man"}
	//Printf 格式化输出
	fmt.Printf("% + v\n", p)     //格式化输出结构
	fmt.Printf("%#v\n", p)       //输出值的 Go 语言表示方法
	fmt.Printf("%T\n", p)        //输出值的类型的 Go 语言表示
	fmt.Printf("%t\n", true)     //输出值的 true 或 false
	fmt.Printf("%b\n", 1024)     //二进制表示
	fmt.Printf("%c\n", 11111111) //数值对应的 Unicode 编码字符
	fmt.Printf("%d\n", 10)       //十进制表示
	fmt.Printf("%o\n", 8)        //八进制表示
	fmt.Printf("%q\n", 22)       //转化为十六进制并附上单引号
	fmt.Printf("%x\n", 1223)     //十六进制表示，用a-f表示
	fmt.Printf("%X\n", 1223)     //十六进制表示，用A-F表示
	fmt.Printf("%U\n", 1233)     //Unicode表示
	fmt.Printf("%b\n", 12.34)    //无小数部分，两位指数的科学计数法6946802425218990p-49
	fmt.Printf("%e\n", 12.345)   //科学计数法，e表示
	fmt.Printf("%E\n", 12.34455) //科学计数法，E表示
	fmt.Printf("%f\n", 12.3456)  //有小数部分，无指数部分
	fmt.Printf("%g\n", 12.3456)  //根据实际情况采用%e或%f输出
	fmt.Printf("%G\n", 12.3456)  //根据实际情况采用%E或%f输出
	fmt.Printf("%s\n", "wqdew")  //直接输出字符串或者[]byte
	fmt.Printf("%q\n", "dedede") //双引号括起来的字符串
	fmt.Printf("%x\n", "abczxc") //每个字节用两字节十六进制表示，a-f表示
	fmt.Printf("%X\n", "asdzxc") //每个字节用两字节十六进制表示，A-F表示
	fmt.Printf("%p\n", 0x123)    //0x开头的十六进制数表示
}

func testPrintFunctions()  {
	// 向io写数据, io可以是任意实现io.Writer的组件，控制台/文件io/网络io等
	fmt.Fprint(os.Stdout, "test")
	// 底层调用的是 Fprint, 默认向控制台打印输出 os.Stdout，
	fmt.Print("this is os.stdout")
	// 以字符串的方式返回,可以组装不同类型的数据
	str := fmt.Sprint("她的语文是", 88, "分,", "英语是", 92.5, "分")
	// 加 ln 是换行
	fmt.Println(str)
	// 加 f 是进行格式化输出
	fmt.Printf("这箱苹果有 %d 斤", 20)
}

func testScanFunctions()  {
	// 从io中读取数据，io可以是任意实现io.Writer的组件，控制台/文件io/网络io等
	data := ""
	fmt.Fscan(os.Stdin, &data)
	fmt.Print(data)
	// 底层调用的是 Fscan，默认使用控制台输入 os.Stdin
	fmt.Scan(&data)
	fmt.Print(data)
	// 将字符串作为输入，字符串作为输出
	fmt.Sscan("这个功能很好用", &data)
	fmt.Print(data)
}

func testErrorFunctions()  {
	// 输出一个error类型
	var name, id = "bueller", 17
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	fmt.Println(err.Error())
}

func main()  {
	testErrorFunctions()
}

