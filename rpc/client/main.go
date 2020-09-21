package main

import (
	"fmt"
	"github.com/fishdemon/go-yehua/rpc/example"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"net/rpc"
	"os"
)

func generateProtoMsg()  {
	// 生成 protobuf msg
	reps := []int64{1,2,3}
	msg := &example.Test{Label: "allen", Type: 29, Reps: reps}

	path := string("./text.txt")
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("error!")
		return
	}

	defer f.Close()

	buffer, err := proto.Marshal(msg)
	f.Write(buffer)
}

func readProtoMsg()  {
	// 读取 protobuf msg
	path := string("./text.txt")
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		return
	}

	defer file.Close()
	fi, err := file.Stat()
	CheckError(err)
	buffer := make([]byte, fi.Size())
	_, err = io.ReadFull(file, buffer)
	CheckError(err)

	msg := &example.Test{}
	err = proto.Unmarshal(buffer, msg)
	CheckError(err)

	fmt.Printf("read: %s\n", msg.String())
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)

}
