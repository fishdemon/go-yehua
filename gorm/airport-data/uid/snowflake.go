package uid

import (
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func init() {
	// Create a new Node with a Node number of 1
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		return
	}
}

func Generate() int64 {
	return node.Generate().Int64()
}
