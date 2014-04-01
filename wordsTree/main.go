package wordsTree

import (
	"io"
	"bytes"
	"strings"
)

const (
	sep_key = ':'
	sep_val = '|'
)

func Create() *Node {
	return &Node{make(map[rune]*Node), nil}
}

func Traverse(node *Node, callback func(key string, values []string)) {
	traverse(node, "", callback)
}

func traverse(node *Node, key string, callback func(key string, values []string)) {
	if node != nil {
		callback(key, node.Values)
	}

	//var n *Node
	for r, n := range node.Children {
		//n = node.Children[r]
		//key += string(r)
		traverse(n, key + string(r), callback)
	}
}

func Serialize(node *Node, writer io.Writer) {
	Traverse(node, serialize_callback(writer))
}

func serialize_callback(w io.Writer) func(key string, values []string) {
	return func(key string, values []string) {
		if len(values) > 0 {
			w.Write([]byte(key))
			w.Write([]byte{sep_key})
			for i, value := range values {
				w.Write([]byte(value))
				if i != len(values) - 1 {
					w.Write([]byte{sep_val})
				}
			}
			w.Write([]byte{'\n'})
		}
	}
}

func Unserialize(r io.Reader) (*Node, error) {
	node := Create()

	buf := new(bytes.Buffer)
	buf.ReadFrom(r)

	lines := strings.SplitN(buf.String(), "\n", -1)
	for _, line := range lines {
		if len(line) > 0 {
			slice := strings.Split(line, ":")
			for _, value := range strings.SplitN(slice[1], "|", -1) {
				node.Insert(slice[0], value)
			}
		}
	}

	return node, nil
	panic("dev")
}
