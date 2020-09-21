package main

import (
	"fmt"
)

func printResult()  {
	red, yellow, blue := "red", "yellow", "blue"

	// 普通数组
	colors := [3]string{red,yellow,blue}
	fmt.Println(colors)			// [red yellow blue]

	// 数组指针
	var ap *[3]string = &colors
	fmt.Println(ap)				// &[red yellow blue]

	// 指针数组
	pa := [3]*string{&red, &yellow,&blue}
	fmt.Println(pa)				// [0xc0001021e0 0xc0001021f0 0xc000102200]
}

func modifyElements()  {
	red, yellow, blue := "red", "yellow", "blue"

	fmt.Println("----------------修改前各数组的值---------------")
	// 普通数组
	colors := [3]string{red,yellow,blue}
	fmt.Println(colors)			// [red yellow blue]
	// 数组指针
	var ap *[3]string = &colors
	fmt.Println(ap)				// &[red yellow blue]
	// 指针数组
	pa := [3]*string{&red, &yellow,&blue}
	fmt.Println(pa)				// [0xc0001021e0 0xc0001021f0 0xc000102200]
	fmt.Println(*pa[0], *pa[1], *pa[2])

	red, yellow, blue = "1", "2", "3"
	// 修改后 colors 中的值
	fmt.Println("----------------修改后各数组的值---------------")
	fmt.Println(colors)
	// 修改后数组指针的值
	fmt.Println(ap)
	// 修改后指针数组的值及内部元素的值
	fmt.Println(pa)
	fmt.Println(*pa[0], *pa[1], *pa[2])
}

func paramPassing()  {
	red, yellow, blue := "red", "yellow", "blue"

	// 普通数组
	colors := [3]string{red,yellow,blue}
	// 数组指针
	var ap *[3]string = &colors
	// 指针数组
	pa := [3]*string{&red, &yellow,&blue}


	// 传递的是数组的拷贝，不会修改原数组
	fmt.Println(colors)			// [red yellow blue]
	fmt.Println("普通数组传递前地址：" + fmt.Sprintf("%p", &colors))
	changeColors(colors)
	fmt.Println(colors)
	fmt.Println()

	// 传递的是数组的引用，方法中的操作会修改原数组
	fmt.Println(ap)				// &[red yellow blue]
	fmt.Println("数组指针传递前地址：" + fmt.Sprintf("%p", ap))
	changeAP(ap)
	fmt.Println(ap)
	fmt.Println()

	// 传递的是数组的拷贝，但是数组中的元素是原元素的引用
	fmt.Println(pa)				// [0xc0001021e0 0xc0001021f0 0xc000102200]
	fmt.Println(*pa[0], *pa[1], *pa[2])
	fmt.Println("指针数组传递前地址：" + fmt.Sprintf("%p", &pa))
	changePA(pa)
	fmt.Println(pa)
	fmt.Println(*pa[0], *pa[1], *pa[2])
}

func changeColors(colors [3]string)  {
	fmt.Println("普通数组传递后地址：" + fmt.Sprintf("%p", &colors))
	colors[0], colors[1], colors[2] = "1", "2", "3"
}

func changeAP(ap *[3]string)  {
	fmt.Println("数组指针传递后地址：" + fmt.Sprintf("%p", ap))

	// 这里注意数组指针的使用方式，采用 (*ap)[0]
	// 若使用 *ap[0] 会报错，它会认为是 *(ap[0])
	(*ap)[0], (*ap)[1], (*ap)[2] = "1", "2", "3"
}

func changePA(pa [3]*string)  {
	fmt.Println("指针数组传递后地址：" + fmt.Sprintf("%p", &pa))
	*pa[0], *pa[1], *pa[2]  = "1", "2","3"
}

func main() {
	//printResult()
	//modifyElements()
	//paramPassing()
}
