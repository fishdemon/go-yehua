package scopy

import (
	"fmt"
	"testing"
)

// go test -run TestAdd -v
func TestCopy(t *testing.T) {
	src := 1
	var dest int
	copy(&dest, src)
	fmt.Println(dest)
}