package wordsTree

import (
	"io"
)

func Create() *Node {
	return &Node{make(map[rune]*Node), nil}
}

func Traverse(node *Node, callback func(key string, values []interface{})) {
	traverse(node, "", callback)
}

func traverse(node *Node, key string, callback func(key string, values []interface{})) {
	var n *Node

	for r := range node.Children {
		n = node.Children[r]
		key += string(r)
		traverse(n, key, callback)
	}

	if node != nil {
		callback(key, node.Values)
	}
}

func Serialize(node *Node, writer *io.Writer) ([]byte, error) {
	cb := func(writer *io.Writer) func(key string, values []interface{}) {
		return func(key string, values []interface{}) {
			k := toBytes(key)
			for _, value := range values {
				v := toBytes(value)
			}
			// write to writer
		}
	}
	Traverse(node, cb(writer))
	panic("dev")
}

func toBytes(x interface{}) []byte {
	panic("dev")
}

func Unserialize(node []byte) (*Node, error) {
	panic("dev")
}

func fromBytes(x []byte) interface{} {
	panic("dev")
}
