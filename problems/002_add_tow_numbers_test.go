package problems

import (
	"reflect"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	q := 0
	fl := &ListNode{
		Val:  0,
		Next: nil,
	}

	cr := fl
	cl1 := l1
	cl2 := l2

	for cl1 != nil || cl2 != nil {
		x, y := 0, 0
		if cl1 != nil {
			x = cl1.Val
			cl1 = cl1.Next
		}

		if cl2 != nil {
			y = cl2.Val
			cl2 = cl2.Next
		}
		s := (x + y + q)
		cr.Next = &ListNode{
			Val:  s % 10,
			Next: nil,
		}

		q = s / 10
		cr = cr.Next
	}

	if q != 0 {
		cr.Next = &ListNode{
			Val:  q,
			Next: nil,
		}
	}
	return fl.Next
}

func TestAddTowNumbers(t *testing.T) {
	l1 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 8,
			Next: &ListNode{
				Val: 7,
			},
		},
	}
	l2 := &ListNode{
		Val: 0,
	}

	res := addTwoNumbers(l1, l2)
	expect := []int{9, 8, 7}
	actual := make([]int, 0, 3)
	for cur := res; cur.Next != nil; {
		actual = append(actual, cur.Val)
		cur = cur.Next
	}
	if reflect.DeepEqual(expect, actual) {
		t.Errorf("actual is %v", actual)
	}
}
