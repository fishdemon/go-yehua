package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
)

func main() {
	// 获取用户目录
	var userHome string
	user, err := user.Current()
	if err == nil {
		userHome = user.HomeDir
		fmt.Println(user.Gid + " " + user.Uid + " " + user.Name + " " + user.Username + " " + user.HomeDir)
		groups, err := user.GroupIds()
		if err == nil {
			for _, item := range groups {
				fmt.Print(item + " ")
			}
		}
		fmt.Println()
	}
	// 获取用户目录
	fmt.Println(os.UserHomeDir())
	fmt.Println(os.UserCacheDir())
	fmt.Println(os.UserConfigDir())
	// 临时目录
	fmt.Println(os.TempDir())

	// 创建目录，并设置权限为 777
	//err1 := os.Mkdir(userHome + "/demo", os.ModePerm)
	//if err1 != nil {
	//	log.Fatal(err1)
	//}

	// 创建多级目录
	//err2 := os.MkdirAll(userHome + "/demo/teacher", os.ModePerm)
	//if err2 != nil {
	//	log.Fatal(err2)
	//}

	// 延时 1 秒
	//time.Sleep(time.Second * 1)

	// 删除当前目录
	// 若存在任何文件或子目录，则会删除失败
	//err3 := os.Remove(userHome + "/demo")
	//if err3 != nil {
	//	log.Fatal(err3)
	//}

	// 删除当前及其所有子目录
	//err4 := os.RemoveAll(userHome + "/demo/teacher")
	//if err4 != nil {
	//	log.Fatal(err4)
	//}

	// 创建一个新文件, 权限默认为 0666
	//file1, err5 := os.Create(userHome + "/demo/info.txt")
	//defer file1.Close()
	//if err5 == nil {
	//	log.Fatal(err5)
	//}

	// 以只读的方式打开文件
	//file2, err6 := os.Open(userHome + "/demo/info.txt")
	//defer file2.Close()
	//if err6 != nil {
	//	log.Fatal(err6)
	//}
	// 文件名是包含路径的
	//fileName := file2.Name()
	//fmt.Println(fileName)

	// 以指定方式打开文件，只读，只写，读写，后面还要加权限
	file3, err7 := os.OpenFile(userHome + "/demo/info.txt", os.O_RDWR, 0666)
	defer file3.Close()
	if err7 != nil {
		log.Fatal(err7)
	}
	// 每次打开文件，重新写入的数据会覆盖之前的数据，因为文件指针偏移量会重置
	// 若一次打开，调用多次 write 会追加, 因为文件指针偏移量会叠加
	// 直接写入string
	file3.WriteString("i hate go \n")
	// 写入字节
	file3.Write([]byte("i hate go very much \n"))
	// 在指定的位置写入字节
	file3.WriteAt([]byte("i hate go very very much \n"), 100)

	// 读取数据
	// 将文件指针偏移量重置
	file3.Seek(0, 0)
	bytes := make([]byte, 1000)
	file3.Read(bytes)
	fmt.Println(string(bytes))

	// 循环读取数据
	file3.Seek(0, 0)
	data := []byte{};
	tmp := make([]byte, 16)
	for n,_ := file3.Read(tmp); n > 0; n,_ = file3.Read(tmp) {
		for i := 0; i <len(tmp) ; i++ {
			data = append(data, tmp[i])
		}
	}

	fmt.Println(string(data))

	// 删除文件
	os.Remove(userHome + "/demo/info.txt")


}


