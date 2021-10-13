package list

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	var num = 1
	node := NewNode(num)

	if v, ok := node.Val.(int); !ok || v != num {
		t.Errorf("unexpected val: %+v\n", node)
	}
}

func TestNewNodeFromArray(t *testing.T) {
	nums := make([]interface{}, 8)
	for i := 0; i < 8; i++ {
		nums[i] = i
	}

	list := NewListFromArray(nums)

	index := 0
	for list.Next != nil {
		if v, ok := list.Val.(int); !ok || index > len(nums)-1 || v != nums[index] {
			t.Errorf("unexpected value: %+v\n", list)
		}
		index++
		list = list.Next
	}
}

func TestNode_InsertHead(t *testing.T) {
	var nilN *Node
	nilN = nilN.InsertHead("val")
	if val, ok := nilN.Val.(string); !ok || val != "val" {
		t.Errorf("unexpected node: %+v\n", nilN)
	}

	node := NewNode("head")
	node.Next = NewNode("tail")
	node = node.InsertHead("new head")
	if val, ok := node.Val.(string); !ok || val != "new head" {
		t.Errorf("unexpected node: %+v\n", node)
	}
}

func TestNode_InsertTail(t *testing.T) {
	var nilN *Node
	nilN = nilN.InsertHead("val")
	if val, ok := nilN.Val.(string); !ok || val != "val" {
		t.Errorf("unexpected node: %+v\n", nilN)
	}

	node := NewNode("head")
	node = node.InsertTail("tail")
	if node == nil || node.Next == nil {
		t.Errorf("unexpected node:  %+v\n", node)
	}
	if val, ok := node.Val.(string); !ok || val != "head" {
		t.Errorf("should be: Node(head)->Node(tail), but: %+v\n", node)
	}
	if val, ok := node.Next.Val.(string); !ok || val != "tail" {
		t.Errorf("unexpected node: %+v\n", node)
	}
}

func TestNode_DeleteHead(t *testing.T) {
	var nilNode *Node
	nilNode, val := nilNode.DeleteHead()
	if nilNode != nil || val != nil {
		t.Errorf("unexpected val: %+v, %+v\n", nilNode, val)
	}

	node := NewNode("head")
	node, val = node.DeleteHead()
	if node != nil || val.(string) != "head" {
		t.Errorf("unexpected val: %+v, %+v\n", node, val)
	}

	node = NewNode("head")
	node.Next = NewNode("tail")
	node, val = node.DeleteHead()
	if node == nil || val.(string) != "head" {
		t.Errorf("unexpected val: %+v, %+v\n", node, val)
	}
}

func TestNode_DeleteTail(t *testing.T) {
	var nilNode *Node
	nilNode, val := nilNode.DeleteTail()
	if nilNode != nil || val != nil {
		t.Errorf("unexpected val: %+v, %+v\n", nilNode, val)
	}

	node := NewNode("head")
	node, val = node.DeleteTail()
	if node != nil || val.(string) != "head" {
		t.Errorf("unexpected val: %+v, %+v\n", node, val)
	}

	node = NewNode("head")
	node.Next = NewNode("tail")
	node, val = node.DeleteTail()
	if node == nil || val.(string) != "tail" {
		t.Errorf("unexpected val: %+v, %+v\n", node, val)
	}
}

func TestNode_Reverse(t *testing.T) {
	nums := make([]interface{}, 8)
	for i := 0; i < 8; i++ {
		nums[i] = i
	}
	list := NewListFromArray(nums)
	list = list.Reverse()
	index := len(nums) - 1
	for list.Next != nil {
		if v, ok := list.Val.(int); !ok || index < 0 || v != nums[index] {
			t.Errorf("unexpected value: %+v\n", list)
		}
		index--
		list = list.Next
	}
}

func TestNode_GetTail(t *testing.T) {
	var nilNode *Node
	if tail := nilNode.GetTail(); tail != nil {
		t.Errorf("want: %+v, but: %+v\n", nil, tail)
	}

	node := NewNode("head")
	if tail := node.GetTail(); tail.Val.(string) != "head" {
		t.Errorf("unexpected tail node: %+v\n", tail)
	}

	node.Next = NewNode("tail")
	if tail := node.GetTail(); tail.Val.(string) != "tail" {
		t.Errorf("unexpected tail node: %+v\n", tail)
	}
}
