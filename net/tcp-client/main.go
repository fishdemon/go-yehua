package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:7575")
	if err != nil {
		fmt.Println("Error dialing, ", err.Error())
		return
	}

	keyboardReader := bufio.NewReader(os.Stdin)
	fmt.Println("first, what is your name?")

	clientName, _ := keyboardReader.ReadString('\n')
	clientName = strings.Trim(clientName, "\n")

	for {
		fmt.Println("what do you want to send?")
		data, _ := keyboardReader.ReadString('\n')
		data = strings.Trim(data, "\n")
		if data == "q" {
			return
		}

		_, err = conn.Write([]byte(clientName + ":" + data))

		b := make([]byte, 1024)
		n, err:= conn.Read(b)
		if err == nil {
			fmt.Println("received from server : ", n, " data: ", string(b))
		}
	}

}


