package model

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	node := NewNode(10)
	if node.Value != 10 {
		t.Errorf("Expected 1, got %d", node.Value)
	}
	if node.Next != -1 {
		t.Errorf("Expected -1, got %d", node.Next)
	}
	if node.Prev != -1 {
		t.Errorf("Expected -1, got %d", node.Prev)
	}
}


func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList[int](10)

	if len(ll.Nodes) != 0 {
		t.Errorf("Expected 0, got %d", len(ll.Nodes))
	}
	if ll.HeadIndex != -1 {
		t.Errorf("Expected -1, got %d", ll.HeadIndex)
	}
	if ll.TailIndex != -1 {
		t.Errorf("Expected -1, got %d", ll.TailIndex)
	}

	node := NewNode(10)
	ll.add(*node)

	if len(ll.Nodes) != 1 {
		t.Errorf("Expected 1, got %d", len(ll.Nodes))
	}

	if ll.HeadIndex != 0 {
		t.Errorf("Expected 0, got %d", ll.HeadIndex)
	}

	if ll.TailIndex != 0 {
		t.Errorf("Expected 0, got %d", ll.TailIndex)
	}

	node2 := NewNode(20)
	ll.add(*node2)

	if len(ll.Nodes) != 2 {
		t.Errorf("Expected 2, got %d", len(ll.Nodes))
	}

	if ll.HeadIndex != 0 {
		t.Errorf("Expected 0, got %d", ll.HeadIndex)
	}

	if ll.TailIndex != 1 {
		t.Errorf("Expected 1, got %d", ll.TailIndex)
	}

	head := ll.Head()
	
	if head.Value != 10 {
		t.Errorf("Expected 10, got %d", head.Value)
	}

	if head.Next != 1 {
		t.Errorf("Expected 1, got %d", head.Next)
	}

	if head.Prev != -1 {
		t.Errorf("Expected -1, got %d", head.Prev)
	}

	tail := ll.Tail()

	if tail.Value != 20 {
		t.Errorf("Expected 20, got %d", tail.Value)
	}

	if tail.Next != -1 {
		t.Errorf("Expected -1, got %d", tail.Next)
	}

	if tail.Prev != 0 {
		t.Errorf("Expected 0, got %d", tail.Prev)
	}
}
