package queue

import (
	"fmt"
	"reflect"
	"testing"
)

// 构建二叉树
var (
	head1 = &TreeNode{
		value: 3,
		left: &TreeNode{
			value: 9,
			left:  nil,
			right: nil,
		},
		right: &TreeNode{
			value: 8,
			left: &TreeNode{
				value: 6,
				left:  nil,
				right: nil,
			},
			right: &TreeNode{
				value: 7,
				left:  nil,
				right: nil,
			},
		},
	}
	expect1 = [][]int{
		[]int{3},
		[]int{9, 8},
		[]int{6, 7},
	}
)

func TestPrintBinaryTree(t *testing.T) {

	res := PrintBinaryTree(head1)
	fmt.Println(res)
	if !reflect.DeepEqual(res, expect1) {
		t.Errorf("expect is %v, but actually get is %v", expect1, res)
	}
}
