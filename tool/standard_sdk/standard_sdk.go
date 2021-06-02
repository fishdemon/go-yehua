package standard_sdk

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"strings"
)

type Node struct {
	Id   int64
	Name string
	Path string
	DescEN string
	Content string
	Parent *Node
	Children []*Node
	GoFiles  []*Node
	Functions []*Node
	Types    []*Node
	Interfaces []*Node
}

// 获取目录
func GetDirectoryTree() *Node {
	root := &Node{
		Path: "/usr/local/go/src",
	}
	transferDirectory(root)
	return root
}

func transferDirectory(node *Node) {
	if  node.Path == "" {
		return
	}

	files, _ := ioutil.ReadDir(node.Path)
	for _, file := range files {
		// 增加子目录
		if file.IsDir() {
			child := &Node{
				Id:       0,
				Name:     file.Name(),
				Path:     node.Path + "/" + file.Name(),
				DescEN:   "",
				Content:  goDoc(file.Name()),
				Parent:   node,
			}
			node.Children = append(node.Children, child)
			continue
		}
		// 增加go文件
		GoFile := &Node{
			Id:       0,
			Name:     file.Name(),
			Path:     node.Path + "/" + file.Name(),
		}
		node.GoFiles = append(node.GoFiles, GoFile)
	}

	// 遍历子节点
	for _, child := range node.Children {
		transferDirectory(child)
		fmt.Println("已完成：" + child.Path)
	}

	// 读取 doc
	if node.Name == "" {
		return
	}

	doc := goDoc(node.Name)
	node.Functions = parseDoc(node.Name, doc)
}

func goDoc(path string) string {
	cmdString := "/usr/local/go/bin/go doc " + path
	res, _ := ExecShell(cmdString)
	return res
}

func parseDoc(packet string, doc string) []*Node {
	nodes := []*Node{}
	funcStrs := strings.Split(doc, "/n")
	for _, funcStr := range funcStrs {
		if strings.HasPrefix(funcStr, "func") {
			temp := funcStr[5:]

			if strings.HasPrefix(temp, "(") {
				index := strings.Index(temp, ")")
				temp = temp[index + 2:]
			}

			index := strings.Index(temp, "(")
			funcName := funcStr[:index]

			node := &Node{
				Id:        0,
				Name:      funcName,
				Path:      packet + "." +funcName,
				DescEN:    "",
				Content:   goDoc(packet + "." + funcName),
			}
			nodes = append(nodes, node)
		}
	}
	return nodes
}

//阻塞式的执行外部shell命令的函数,等待执行完毕并返回标准输出
func ExecShell(s string) (string, error){
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", s)

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out

	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()

	return out.String(), err
}

func ExecCommand(commandName string, params []string) bool {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	fmt.Println(cmd.Args)
	//StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return false
	}

	cmd.Start()
	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}

	//阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的
	cmd.Wait()
	return true
}