package wordsTree

const capacity = 10
const REMOVE_ALL_VALUES = "delete_all_values:true"

type Node struct {
	Children map[rune]*Node
	Values   []interface{}
}

func (n *Node) Find(key string) (interface{}, bool) {
	node := n
	found := false

	for _, char := range key {
		if node, found = node.Children[char]; !found {
			return nil, false
		}
	}

	return node.Values, true
}

func (n *Node) Insert(key string, value interface{}) bool {
	node := n

	for _, char := range key {
		if temp, found := node.Children[char]; !found {
			node.Children[char] = &Node{make(map[rune]*Node), make([]interface{}, 0, capacity)}
			node = node.Children[char]
		} else {
			node = temp
		}
	}

	for _, v := range node.Values {
		if v == value {
			return false
		}
	}

	node.Values = append(node.Values, value)
	return true
}

func (n *Node) Update(key string, oldValue, newValue interface{}) bool {
	node := n
	var found bool

	for _, char := range key {
		if node, found = node.Children[char]; !found {
			return false
		}
	}

	for i, v := range node.Values {
		if v == oldValue {
			node.Values[i] = newValue
			return true
		}
	}

	return false
}

func (n *Node) Remove(key string, value interface{}) bool {
	var r rune
	var found bool
	var last *Node
	node := n

	for i, char := range key {
		r = char
		if last, found = node.Children[char]; !found {
			return false
		}
		if len([]rune(key))-1 != i {
			node = last
		}
	}

	if value == REMOVE_ALL_VALUES {
		if len(last.Children) == 0 {
			delete(node.Children, r)
		} else {
			last.Values = make([]interface{}, 0, capacity)
		}
		return true
	} else {
		for i, v := range last.Values {
			if v == value {
				last.Values = append(last.Values[:i], last.Values[i+1:]...)
				return true
			}
		}
	}

	return false
}
