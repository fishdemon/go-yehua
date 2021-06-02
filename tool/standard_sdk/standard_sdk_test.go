package standard_sdk

import (
	"fmt"
	"testing"
)

func TestGetDirectoryTree(t *testing.T) {
	GetDirectoryTree()
}

func TestExec_shell(t *testing.T) {
	res, _ := ExecShell("go doc io")
	fmt.Println(res)
}

