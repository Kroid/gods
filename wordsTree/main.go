package wordsTree

import (
	"io"
	"strconv"
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

func Serialize(node *Node, writer io.Writer) {
	cb := func(writer io.Writer) func(key string, values []interface{}) {
		return func(key string, values []interface{}) {
			writer.Write(toBytes(key))
			writer.Write([]byte{':'})
			for _, value := range values {
				writer.Write(toBytes(value))
				writer.Write([]byte{'|'})
			}
			writer.Write([]byte{'\n'})
		}
	}
	Traverse(node, cb(writer))
	//panic("dev")
}

func toBytes(x interface{}) []byte {
	switch x.(type){
	case bool:
		return []byte(strconv.FormatBool(x.(bool)))
	case int:
		return []byte(strconv.FormatInt(int64(x.(int)), 10))
	case string:
		return []byte(x.(string))
	}
	panic("dev")
}

func Unserialize(node []byte) (*Node, error) {
	panic("dev")
}

func fromBytes(x []byte) interface{} {
	panic("dev")
}
