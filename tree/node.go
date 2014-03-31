package tree

type Node struct {
	Left, Right, Parent *Node
	Key, Value          interface{}
	Less                func(a, b interface{}) bool
}

func (n *Node) Enumerate(callback func(key, value interface{})) {
	if n.Left != nil {
		n.Left.Enumerate(callback)
	}

	callback(n.Key, n.Value)

	if n.Right != nil {
		n.Right.Enumerate(callback)
	}
}

func (n *Node) Insert(key, value interface{}) bool {
	if key == n.Key {
		return false
	}

	if n.Key == nil {
		n.Key, n.Value = key, value
		return true
	}

	if n.Less(key, n.Key) {
		if n.Left == nil {
			n.Left = &Node{nil, nil, n, key, value, n.Less}
			return true
		}
		return n.Left.Insert(key, value)
	}

	if n.Right == nil {
		n.Right = &Node{nil, nil, n, key, value, n.Less}
		return true
	}

	return n.Right.Insert(key, value)
}

func (n *Node) Update(key, value interface{}) bool {
	if key == n.Key {
		n.Value = value
		return true
	}

	if n.Less(key, n.Key) {
		if n.Left != nil {
			return n.Left.Update(key, value)
		}
	} else {
		if n.Right != nil {
			return n.Right.Update(key, value)
		}
	}
	return false
}

func (n *Node) Find(key interface{}) (interface{}, bool) {
	if key == n.Key {
		return n.Value, true
	}

	if n.Less(key, n.Key) {
		if n.Left != nil {
			return n.Left.Find(key)
		}
	} else {
		if n.Right != nil {
			return n.Right.Find(key)
		}
	}
	return nil, false
}

func (n *Node) Remove(key interface{}) bool {
	if n.Less(key, n.Key) {
		if n.Left != nil {
			return n.Left.Remove(key)
		}
	} else if key == n.Key {
		return removeEqual(n, key)
	} else {
		if n.Right != nil {
			return n.Right.Remove(key)
		}
	}

	return false
}

func removeEqual(n *Node, key interface{}) bool {
	if n.Left == nil && n.Right == nil {
		if n.Parent.Left == n {
			n.Parent.Left = nil
		} else {
			n.Parent.Right = nil
		}
		return true
	}

	if n.Left == nil {
		n.Right.Parent = n.Parent
		if n.Parent.Left == n {
			n.Parent.Left = n.Right
		} else {
			n.Parent.Right = n.Right
		}
		return true
	}

	if n.Right == nil {
		n.Left.Parent = n.Parent
		if n.Parent.Left == n {
			n.Parent.Left = n.Left
		} else {
			n.Parent.Right = n.Left
		}
		return true
	}

	// hard case
	min := n.Right
	for {
		if min.Left == nil {
			break
		}
		min = min.Left
	}
	n.Right.Left = nil

	if min.Right != nil {
		right := min.Right
		for {
			if right.Right == nil {
				break
			}
			right = right.Right
		}
		right.Right = n.Right
		n.Right.Parent = right
		min.Right.Parent = n
	} else {
		min.Right = n.Right
	}

	n.Key, n.Value = min.Key, min.Value
	n.Right = min.Right

	return true

}