package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	var s = new(Stack)
	s.Push(3)
	if s.data[0] != 3 {
		t.Log("3 should be on stack!")
		t.Fail()
	}
}


func TestPop(t *testing.T) {
	var s = new(Stack)
	s.Push(3)
	v := s.Pop()
	if 3 != v {
		t.Log("Value popped off stack should be 3")
		t.Fail()
	}
	
}
