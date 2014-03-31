package tree

func Create(less func(a, b interface{}) bool) *Node {
	return &Node{nil, nil, nil, nil, nil, less}
}

func Serialize(t *Node) ([]byte, error) {
	panic("dev")
}
func Unserialize(t []byte) (*Node, error) {
	panic("dev")
}
