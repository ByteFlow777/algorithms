package list

type Node struct {
	Next *Node
	Val  interface{}
}

// NewNode return a new LinkedList node
func NewNode(val interface{}) *Node {
	return &Node{
		Next: nil,
		Val:  val,
	}
}

// NewListFromArray use the given values to build a new LinkedList,
// If given array is nil return nil.
func NewListFromArray(values []interface{}) *Node {
	if len(values) <= 0 {
		return nil
	}

	head := NewNode(values[0])
	dummy := head
	for i, n := 1, len(values); i < n; i++ {
		dummy.Next = NewNode(values[i])
		dummy = dummy.Next
	}

	return head
}

// InsertHead insert a new node into the header.
// You should assign the result to the old Node.
func (n *Node) InsertHead(val interface{}) *Node {
	newHead := NewNode(val)
	newHead.Next = n
	return newHead
}

// InsertTail insert a new node to the end.
// You should assign the result to the old Node.
func (n *Node) InsertTail(val interface{}) *Node {
	if n == nil {
		return NewNode(val)
	}

	dummy := n
	for dummy.Next != nil {
		dummy = dummy.Next
	}
	dummy.Next = NewNode(val)
	return n
}

// DeleteHead remove the first node and return its value.
// If given node is nil return nil.
func (n *Node) DeleteHead() (node *Node, val interface{}) {
	if n == nil {
		return nil, nil
	}
	node = n.Next

	// unlink header and it's Next node
	n.Next = nil
	val = n.Val
	return
}

// DeleteTail remove the last node and return its value.
func (n *Node) DeleteTail() (*Node, interface{}) {
	if n == nil {
		return nil, nil
	}
	if n.Next == nil {
		return nil, n.Val
	}
	dummy := n
	for dummy.Next != nil && dummy.Next.Next != nil {
		dummy = dummy.Next
	}
	val := dummy.Next.Val
	dummy.Next.Next = nil
	return n, val
}

// GetTail return the last node.
// If head is nil return nil.
// If head does not have Next node, return itself.
func (n *Node) GetTail() *Node {
	if n == nil {
		return nil
	}
	for n.Next != nil {
		n = n.Next
	}
	return n
}

// Reverse the LinkedList and return a new header.
// You should assign the result to the old Node.
func (n *Node) Reverse() *Node {
	var dummy *Node
	for n != nil {
		temp := n.Next
		n.Next = dummy
		dummy = n
		n = temp
	}
	return dummy
}

// ToArray rebuild LinkedList to array
func (n *Node) ToArray() []interface{} {
	nodes := make([]interface{}, 0)
	for n != nil {
		nodes = append(nodes, n.Val)
	}
	return nodes
}
